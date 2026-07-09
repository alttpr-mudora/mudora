package rom

import (
	"sort"
	"strings"
)

// Entry is the item placed at a known location in a ROM.
type Entry struct {
	Address  uint32
	Location string
	Item     string
}

const (
	pcIsEncrypted = 0x180087
	pcStaticKey   = 0x1800B0
	pcPayload     = 0xEABC
	chestBase     = 0xE96E // item-byte base of the stride-3 table (non-race format)
	chestStride   = 3
	chestCount    = 168
	pcHash        = 0x180215
	hashItemCount = 5
	pcSeedString  = 0x7FC0 // SNES cartridge title field, overwritten with "VT <permalink>"
	seedStringLen = 21
)

// inPlaceRegion is a range of bytes encrypted in-place using XXTEA.
// All addresses in [lo, hi] belong to this region.
type inPlaceRegion struct{ tableBase, lo, hi uint32 }

// Race ROMs encrypt these four table regions in-place using the same XXTEA key.
// Source: z3randomizer tables.asm [encrypted] annotations and heartpieces.asm macro calls.
var encryptedRegions = []inPlaceRegion{
	{0x180000, 0x180000, 0x180006}, // HeartPieceIndoorValues
	{0x180010, 0x180010, 0x180017}, // SpriteItemValues (block 0 only)
	{0x180140, 0x180140, 0x18014A}, // HeartPieceOutdoorValues
	{0x180150, 0x180150, 0x180159}, // HeartContainerBossValues
}

func Inspect(data []byte) []Entry {
	encrypted := len(data) > pcIsEncrypted+1 &&
		(data[pcIsEncrypted] != 0 || data[pcIsEncrypted+1] != 0)

	var key [4]uint32
	var plain [chestCount]byte
	if encrypted {
		key = readKey(data)
		for blk := 0; blk < 21; blk++ {
			off := pcPayload + blk*8
			dec := xxteaDecrypt(data[off:off+8], key)
			copy(plain[blk*8:], dec[:])
		}
	}

	entries := make([]Entry, 0, len(Locations))
	for addr, loc := range Locations {
		item := "UNKNOWN"
		if encrypted && isChestAddr(addr) {
			n := (int(addr) - chestBase) / chestStride
			if n >= 0 && n < chestCount {
				if name, ok := ItemAt(uint32(plain[n])); ok {
					item = name
				}
			}
		} else if encrypted {
			if b, ok := decryptInPlace(data, key, addr); ok {
				if name, ok := ItemAt(uint32(b)); ok {
					item = name
				}
			} else if i := int(addr); i >= 0 && i < len(data) {
				if name, ok := ItemAt(uint32(data[i])); ok {
					item = name
				}
			}
		} else if i := int(addr); i >= 0 && i < len(data) {
			if name, ok := ItemAt(uint32(data[i])); ok {
				item = name
			}
		}
		entries = append(entries, Entry{Address: addr, Location: loc, Item: item})
	}

	for dungeon, prize := range DungeonPrizes(data) {
		loc, ok := prizeLocationName[dungeon]
		if !ok {
			continue
		}
		entries = append(entries, Entry{Address: prizeAddrs[dungeon].category, Location: loc, Item: prize})
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Address < entries[j].Address })
	return entries
}

// Start screen hash symbols uses a separate lookup table than normal items.
// Names match items.txt/icons.go.
var hashSymbols = map[byte]string{
	0: "Bow", 1: "Boomerang", 2: "Hookshot", 3: "Bomb",
	4: "Mushroom", 5: "Powder", 6: "Ice Rod", 7: "Pendant of Courage",
	8: "Bombos", 9: "Ether", 10: "Quake", 11: "Lamp",
	12: "Hammer", 13: "Shovel", 14: "Flute", 15: "Bug Net", 16: "Book of Mudora",
	17: "Bottle", 18: "Green Potion", 19: "Cane of Somaria", 20: "Cape",
	21: "Magic Mirror", 22: "Pegasus Boots", 23: "Progressive Glove", 24: "Flippers",
	25: "Moon Pearl", 26: "Progressive Shield", 27: "Progressive Mail", 28: "Heart",
	29: "Map", 30: "Compass", 31: "Big Key",
}

func GetHash(data []byte) ([]string, bool) {
	hashItems := make([]string, hashItemCount)

	for i := range hashItemCount {
		addr := pcHash + i
		value := data[addr]
		item, ok := hashSymbols[value]
		if !ok {
			return nil, false
		}

		hashItems[i] = item
	}

	return hashItems, true
}

func GetPermalink(data []byte) (string, bool) {
	if pcSeedString+seedStringLen > len(data) {
		return "", false
	}

	raw := strings.TrimSpace(string(data[pcSeedString : pcSeedString+seedStringLen]))
	hash, ok := strings.CutPrefix(raw, "VT ")
	if !ok {
		return "", false
	}

	return hash, true
}

func ReadBytes(data []byte, base int64, amt int64) []byte {
	bytes := make([]byte, amt)

	for offset := range amt {
		address := base + int64(offset)
		value := data[address]

		bytes[offset] = value
	}

	return bytes
}

func isChestAddr(addr uint32) bool {
	return addr >= 0xE96E && addr <= 0xEB63
}

func decryptInPlace(data []byte, key [4]uint32, addr uint32) (byte, bool) {
	for _, r := range encryptedRegions {
		if addr < r.lo || addr > r.hi {
			continue
		}
		offset := addr - r.tableBase
		blockStart := r.tableBase + (offset/8)*8
		if int(blockStart)+8 > len(data) {
			return 0, false
		}
		dec := xxteaDecrypt(data[blockStart:blockStart+8], key)
		return dec[offset%8], true
	}
	return 0, false
}

func readKey(data []byte) [4]uint32 {
	var key [4]uint32
	for i := range key {
		off := pcStaticKey + i*4
		key[i] = uint32(data[off]) | uint32(data[off+1])<<8 | uint32(data[off+2])<<16 | uint32(data[off+3])<<24
	}
	return key
}

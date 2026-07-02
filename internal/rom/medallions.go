package rom

var medallionAddr = map[string]uint32{
	"Misery Mire": 0x180022,
	"Turtle Rock": 0x180023,
}

var medallionByte = map[byte]string{
	0x00: "Bombos",
	0x01: "Ether",
	0x02: "Quake",
}

func RequiredMedallions(data []byte) map[string]string {
	medallions := make(map[string]string, len(medallionAddr))

	for name, addr := range medallionAddr {
		if int(addr) >= len(data) {
			continue
		}

		if medallion, ok := medallionByte[data[addr]]; ok {
			medallions[name] = medallion
		}
	}

	return medallions
}

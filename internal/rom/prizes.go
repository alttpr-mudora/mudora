package rom

type prizeAddr struct {
	category uint32
	detail   uint32
}

var prizeAddrs = map[string]prizeAddr{
	"Eastern Palace":     {0x1209D, 0x180070},
	"Desert Palace":      {0x1209E, 0x180072},
	"Tower of Hera":      {0x120A5, 0x180071},
	"Palace of Darkness": {0x120A1, 0x180073},
	"Swamp Palace":       {0x120A0, 0x180079},
	"Skull Woods":        {0x120A3, 0x180074},
	"Thieves Town":       {0x120A6, 0x180076},
	"Ice Palace":         {0x120A4, 0x180078},
	"Misery Mire":        {0x120A2, 0x180077},
	"Turtle Rock":        {0x120A7, 0x180075},
}

var prizeSignatures = map[[2]byte]string{
	{0x04, 0x69}: "Pendant of Courage",
	{0x01, 0x69}: "Pendant of Wisdom",
	{0x02, 0x69}: "Pendant of Power",
	{0x02, 0x7F}: "Crystal 1",
	{0x10, 0x79}: "Crystal 2",
	{0x40, 0x6C}: "Crystal 3",
	{0x20, 0x6D}: "Crystal 4",
	{0x04, 0x6E}: "Crystal 5",
	{0x01, 0x6F}: "Crystal 6",
	{0x08, 0x7C}: "Crystal 7",
}

func DungeonPrizes(data []byte) map[string]string {
	prizes := make(map[string]string, len(prizeAddrs))

	for dungeon, a := range prizeAddrs {
		if int(a.category) >= len(data) || int(a.detail) >= len(data) {
			continue
		}

		sig := [2]byte{data[a.category], data[a.detail]}
		if name, ok := prizeSignatures[sig]; ok {
			prizes[dungeon] = name
		}
	}

	return prizes
}

package graph

var progressionItems = map[string]bool{
	"Progressive Sword":    true,
	"Master Sword":         true,
	"Progressive Shield":   true,
	"Progressive Armor":    true,
	"Progressive Glove":    true,
	"Progressive Bow":      true,
	"Bow":                  true,
	"Silver Arrow Upgrade": true,
	"Hookshot":             true,
	"Hammer":               true,
	"Fire Rod":             true,
	"Ice Rod":              true,
	"Bombos":               true,
	"Ether":                true,
	"Quake":                true,
	"Lamp":                 true,
	"Flippers":             true,
	"Moon Pearl":           true,
	"Magic Mirror":         true,
	"Pegasus Boots":        true,
	"Cane of Somaria":      true,
	"Cane of Byrna":        true,
	"Cape":                 true,
	"Mushroom":             true,
	"Powder":               true,
	"Ocarina":              true,
	"Flute":                true,
	"Book of Mudora":       true,
	"Bug Net":              true,
	"Half Magic":           true,
	"Quarter Magic":        true,
	"Boomerang":            true,
	"Boomerang (Red)":      true,
	"Shovel":               true,
	"Bombs (10)":           true,

	"Bottle":                true,
	"Bottle (Red Potion)":   true,
	"Bottle (Green Potion)": true,
	"Bottle (Blue Potion)":  true,
	"Bottle (Bee)":          true,
	"Bottle (Super bee)":    true,
	"Bottle (Faerie)":       true,

	"Heart Container": false,
	"Piece of Heart":  false,

	"Pendant of Courage": true,
	"Pendant of Power":   true,
	"Pendant of Wisdom":  true,
	"Crystal 1":          true,
	"Crystal 2":          true,
	"Crystal 3":          true,
	"Crystal 4":          true,
	"Crystal 5":          true,
	"Crystal 6":          true,
	"Crystal 7":          true,

	"Key (Hyrule Castle)": true,
	"Key (Castle Tower)":  true,

	"Big Key (Eastern Palace)": true,

	"Key (Desert Palace)":     true,
	"Big Key (Desert Palace)": true,

	"Key (Tower of Hera)":     true,
	"Big Key (Tower of Hera)": true,

	"Key (Palace of Darkness)":     true,
	"Big Key (Palace of Darkness)": true,

	"Key (Swamp Palace)":     true,
	"Big Key (Swamp Palace)": true,

	"Key (Skull Woods)":     true,
	"Big Key (Skull Woods)": true,

	"Key (Thieves Town)":     true,
	"Big Key (Thieves Town)": true,

	"Key (Ice Palace)":     true,
	"Big Key (Ice Palace)": true,

	"Key (Misery Mire)":     true,
	"Big Key (Misery Mire)": true,

	"Key (Turtle Rock)":     true,
	"Big Key (Turtle Rock)": true,

	"Key (Ganon's Tower)":     true,
	"Big Key (Ganon's Tower)": true,

	"Rescue Zelda":    true,
	"Defeat Agahnim":  true,
	"Defeat Agahnim2": true,
}

func isProgression(item string) bool {
	return progressionItems[item]
}

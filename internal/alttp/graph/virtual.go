package graph

import "github.com/alttpr-mudora/mudora/internal/alttp/logic"

var (
	locationRules   = logic.LocationRules()
	completionRules = logic.CompletionRules()
)

type virtualLocation struct {
	region string
	rule   logic.Rule
	item   string
}

var virtualLocations = map[string]virtualLocation{
	"Castle Tower": {region: "Castle Tower", rule: completionRules["Castle Tower"], item: "Defeat Agahnim"},
	"Ganons Tower": {region: "Ganons Tower", rule: completionRules["Ganons Tower"], item: "Defeat Agahnim2"},
}

var prizeDungeonRegion = map[string]string{
	"Eastern Palace - Prize":     "Eastern Palace",
	"Desert Palace - Prize":      "Desert Palace",
	"Tower of Hera - Prize":      "Tower of Hera",
	"Palace of Darkness - Prize": "Palace of Darkness",
	"Swamp Palace - Prize":       "Swamp Palace",
	"Skull Woods - Prize":        "Skull Woods",
	"Thieves' Town - Prize":      "Thieves Town",
	"Ice Palace - Prize":         "Ice Palace",
	"Misery Mire - Prize":        "Misery Mire",
	"Turtle Rock - Prize":        "Turtle Rock",
}

func withVirtualItems(itemAt map[string]string) map[string]string {
	merged := make(map[string]string, len(itemAt)+len(virtualLocations))
	for loc, item := range itemAt {
		merged[loc] = item
	}
	for loc, v := range virtualLocations {
		merged[loc] = v.item
	}
	return merged
}

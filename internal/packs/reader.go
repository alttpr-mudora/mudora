package packs

import (
	"embed"
)

//go:embed alttpr_emotracker_emosaru/locations/*.json
var emosaruOfficialPack embed.FS

var availablePacks = map[string]embed.FS{
	"emosaru": emosaruOfficialPack,
}

func GetRawRules(name string) (map[string]string, bool) {
	pack, ok := availablePacks[name]

	if !ok {
		return nil, false
	}

	rules, ok := readRules(&pack)

	if !ok {
		return nil, false
	}

	return rules, true
}

func readRules(pack *embed.FS) (map[string]string, bool) {
	rules := make(map[string]string)

	return rules, true
}

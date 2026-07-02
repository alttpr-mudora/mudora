package graph

import "github.com/alttpr-mudora/mudora/internal/alttp/logic"

var regionRules = logic.RegionRules()

func regionAccessibility(items *logic.Items, settings *logic.Settings) map[string]bool {
	access := make(map[string]bool, len(regionRules))
	lookup := func(name string) bool { return access[name] }

	for changed := true; changed; {
		changed = false
		for name, rule := range regionRules {
			if !access[name] && rule(items, settings, lookup) {
				access[name] = true
				changed = true
			}
		}
	}

	return access
}

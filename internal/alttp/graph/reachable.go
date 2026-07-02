package graph

import "github.com/alttpr-mudora/mudora/internal/alttp/logic"

func ruleFor(loc string) (logic.Rule, bool) {
	if rule, ok := locationRules[loc]; ok {
		return rule, true
	}
	if v, ok := virtualLocations[loc]; ok {
		return v.rule, true
	}
	if region, ok := prizeDungeonRegion[loc]; ok {
		return completionRules[region], true
	}
	return nil, false
}

func regionFor(loc string) string {
	if region, ok := logic.LocationRegion[loc]; ok {
		return region
	}
	if v, ok := virtualLocations[loc]; ok {
		return v.region
	}
	return prizeDungeonRegion[loc]
}

func locationReachable(loc string, items *logic.Items, settings *logic.Settings, access map[string]bool) bool {
	if region := regionFor(loc); region != "" && !access[region] {
		return false
	}
	rule, ok := ruleFor(loc)
	if !ok {
		return false
	}
	return rule(items, settings, func(name string) bool { return access[name] })
}

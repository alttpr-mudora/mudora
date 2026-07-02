package logic

// RegionAccess reports whether the named region can be entered. A future
// solver supplies this (memoized fixed-point over accumulated items); it
// lets a region's own entrance rule reference another region's, as several
// PHP region files do (e.g. Desert Palace's entrance checks Mire's).
type RegionAccess func(name string) bool

type Rule func(items *Items, settings *Settings, regions RegionAccess) bool

func alwaysAccessible(items *Items, settings *Settings, regions RegionAccess) bool {
	return true
}

var locationRuleSets = []func() map[string]Rule{
	hyruleCastleEscapeLocationRules,
	easternPalaceLocationRules,
	desertPalaceLocationRules,
	hyruleCastleTowerLocationRules,
	fountainsLocationRules,
	lightWorldLocationRules,
	darkWorldLocationRules,
	deathMountainLocationRules,
	mireLocationRules,
	towerOfHeraLocationRules,
	palaceOfDarknessLocationRules,
	swampPalaceLocationRules,
	skullWoodsLocationRules,
	icePalaceLocationRules,
	miseryMireLocationRules,
	turtleRockLocationRules,
	ganonsTowerLocationRules,
	thievesTownLocationRules,
}

var regionRuleSets = []func() map[string]Rule{
	hyruleCastleEscapeRegionRules,
	easternPalaceRegionRules,
	desertPalaceRegionRules,
	hyruleCastleTowerRegionRules,
	lightWorldRegionRules,
	darkWorldRegionRules,
	deathMountainRegionRules,
	mireRegionRules,
	towerOfHeraRegionRules,
	palaceOfDarknessRegionRules,
	swampPalaceRegionRules,
	skullWoodsRegionRules,
	icePalaceRegionRules,
	miseryMireRegionRules,
	turtleRockRegionRules,
	ganonsTowerRegionRules,
	thievesTownRegionRules,
}

var completionRuleSets = []func() map[string]Rule{
	hyruleCastleEscapeCompletionRules,
	easternPalaceCompletionRules,
	desertPalaceCompletionRules,
	hyruleCastleTowerCompletionRules,
	towerOfHeraCompletionRules,
	palaceOfDarknessCompletionRules,
	swampPalaceCompletionRules,
	skullWoodsCompletionRules,
	icePalaceCompletionRules,
	miseryMireCompletionRules,
	turtleRockCompletionRules,
	ganonsTowerCompletionRules,
	thievesTownCompletionRules,
}

func LocationRules() map[string]Rule {
	return mergeRuleSets(locationRuleSets)
}

// RegionRules returns each region's entrance ("can_enter" in the PHP
// source) rule, keyed by region name.
func RegionRules() map[string]Rule {
	return mergeRuleSets(regionRuleSets)
}

// CompletionRules returns each region's completion ("can_complete" in the
// PHP source) rule, keyed by region name. This is what gates the event
// pseudo-items (e.g. "Defeat Agahnim") a future solver grants once a
// region's boss/prize condition is met.
func CompletionRules() map[string]Rule {
	return mergeRuleSets(completionRuleSets)
}

func mergeRuleSets(sets []func() map[string]Rule) map[string]Rule {
	merged := make(map[string]Rule)

	for _, set := range sets {
		for name, rule := range set() {
			merged[name] = rule
		}
	}

	return merged
}

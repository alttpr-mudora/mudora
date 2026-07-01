package logic

func hyruleCastleEscapeLocationRules() map[string]Rule {
	hasLampOrAdvancedFireRod := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.HasAtLeast("Lamp", settings.LampRequireCount) ||
			(settings.ItemPlacementAdvanced && items.Has("Fire Rod"))
	}

	secretRoom := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.CanLiftRocks() ||
			(hasLampOrAdvancedFireRod(items, settings, regions) &&
				items.Has("Key (Hyrule Castle)") &&
				items.CanKillMostThings(settings, 5))
	}

	needsKeyH2AndCanKillMostThings := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Key (Hyrule Castle)") && items.CanKillMostThings(settings, 5)
	}

	return map[string]Rule{
		"Sewers - Secret Room - Left":     secretRoom,
		"Sewers - Secret Room - Middle":   secretRoom,
		"Sewers - Secret Room - Right":    secretRoom,
		"Sewers - Dark Cross":             hasLampOrAdvancedFireRod,
		"Hyrule Castle - Boomerang Chest": needsKeyH2AndCanKillMostThings,
		"Hyrule Castle - Zelda's Cell":    needsKeyH2AndCanKillMostThings,
		"Sanctuary":                       alwaysAccessible,
		"Hyrule Castle - Map Chest":       alwaysAccessible,
		"Secret Passage":                  alwaysAccessible,
		"Link's Uncle":                    alwaysAccessible,
	}
}

func hyruleCastleEscapeRegionRules() map[string]Rule {
	return map[string]Rule{}
}

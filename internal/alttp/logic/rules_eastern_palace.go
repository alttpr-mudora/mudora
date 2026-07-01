package logic

func easternPalaceLocationRules() map[string]Rule {
	hasLampOrAdvancedFireRod := func(items *Items, settings *Settings) bool {
		return items.HasAtLeast("Lamp", settings.LampRequireCount) ||
			(settings.ItemPlacementAdvanced && items.Has("Fire Rod"))
	}

	return map[string]Rule{
		"Eastern Palace - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Eastern Palace)")
		},
		"Eastern Palace - Big Key Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.HasAtLeast("Lamp", settings.LampRequireCount)
		},
		"Eastern Palace - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanShootArrows(settings, 1) &&
				hasLampOrAdvancedFireRod(items, settings) &&
				items.Has("Big Key (Eastern Palace)") &&
				CanBeatArmosKnights(items, settings)
		},
		"Eastern Palace - Compass Chest":    alwaysAccessible,
		"Eastern Palace - Cannonball Chest": alwaysAccessible,
		"Eastern Palace - Map Chest":        alwaysAccessible,
	}
}

func easternPalaceRegionRules() map[string]Rule {
	return map[string]Rule{
		"Eastern Palace": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Rescue Zelda")
		},
	}
}

func easternPalaceCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Eastern Palace": easternPalaceLocationRules()["Eastern Palace - Boss"],
	}
}

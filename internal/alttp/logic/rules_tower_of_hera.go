package logic

func towerOfHeraMain(items *Items, settings *Settings, regions RegionAccess) bool {
	return (items.Has("Magic Mirror") || (items.Has("Hookshot") && items.Has("Hammer"))) &&
		regions("West Death Mountain")
}

func towerOfHeraLocationRules() map[string]Rule {
	return map[string]Rule{
		"Tower of Hera - Big Key Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanLightTorches() && items.Has("Key (Tower of Hera)")
		},
		"Tower of Hera - Compass Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return towerOfHeraMain(items, settings, regions) && items.Has("Big Key (Tower of Hera)")
		},
		"Tower of Hera - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return towerOfHeraMain(items, settings, regions) && items.Has("Big Key (Tower of Hera)")
		},
		"Tower of Hera - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return towerOfHeraMain(items, settings, regions) &&
				CanBeatMoldorm(items, settings) &&
				items.Has("Big Key (Tower of Hera)")
		},
		"Tower of Hera - Basement Cage": alwaysAccessible,
		"Tower of Hera - Map Chest":     alwaysAccessible,
	}
}

func towerOfHeraRegionRules() map[string]Rule {
	return map[string]Rule{
		"Tower of Hera": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Rescue Zelda") && towerOfHeraMain(items, settings, regions)
		},
	}
}

func towerOfHeraCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Tower of Hera": towerOfHeraLocationRules()["Tower of Hera - Boss"],
	}
}

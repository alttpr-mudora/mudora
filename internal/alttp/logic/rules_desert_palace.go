package logic

func desertPalaceLocationRules() map[string]Rule {
	return map[string]Rule{
		"Desert Palace - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Desert Palace)")
		},
		"Desert Palace - Big Key Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Key (Desert Palace)") && items.CanKillMostThings(settings, 5)
		},
		"Desert Palace - Compass Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Key (Desert Palace)")
		},
		"Desert Palace - Torch": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Pegasus Boots")
		},
		"Desert Palace - Map Chest": alwaysAccessible,
		"Desert Palace - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return desertPalaceCanEnter(items, settings, regions) &&
				(items.CanLiftRocks() ||
					(settings.CanBootsClip && items.Has("Pegasus Boots")) ||
					(settings.CanSuperSpeed && items.CanSpinSpeed()) ||
					settings.CanOneFrameClipOW ||
					(items.Has("Magic Mirror") && regions("Mire"))) &&
				items.CanLightTorches() &&
				items.Has("Big Key (Desert Palace)") && items.Has("Key (Desert Palace)") &&
				CanBeatLanmolas(items, settings)
		},
	}
}

func desertPalaceCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(items.Has("Book of Mudora") ||
			(settings.CanBootsClip && items.Has("Pegasus Boots")) ||
			settings.CanOneFrameClipOW ||
			(items.Has("Magic Mirror") && regions("Mire")))
}

func desertPalaceRegionRules() map[string]Rule {
	return map[string]Rule{
		"Desert Palace": desertPalaceCanEnter,
	}
}

func desertPalaceCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Desert Palace": desertPalaceLocationRules()["Desert Palace - Boss"],
	}
}

package logic

func swampPalaceCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(1)) && items.HasHealth(7) && items.HasABottle())) &&
		items.Has("Flippers") && regions("South Dark World") &&
		items.Has("Moon Pearl") && items.Has("Magic Mirror")
}

func swampPalaceLocationRules() map[string]Rule {
	keyAndHammer := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Key (Swamp Palace)") && items.Has("Hammer")
	}

	return map[string]Rule{
		"Swamp Palace - Entrance": alwaysAccessible,
		"Swamp Palace - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Key (Swamp Palace)") && items.Has("Hammer") && items.Has("Big Key (Swamp Palace)")
		},
		"Swamp Palace - Big Key Chest": keyAndHammer,
		"Swamp Palace - Map Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.Has("Key (Swamp Palace)")
		},
		"Swamp Palace - West Chest":    keyAndHammer,
		"Swamp Palace - Compass Chest": keyAndHammer,
		"Swamp Palace - Flooded Room - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && keyAndHammer(items, settings, regions)
		},
		"Swamp Palace - Flooded Room - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && keyAndHammer(items, settings, regions)
		},
		"Swamp Palace - Waterfall Room": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && keyAndHammer(items, settings, regions)
		},
		"Swamp Palace - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && keyAndHammer(items, settings, regions) && CanBeatArrghus(items, settings)
		},
	}
}

func swampPalaceRegionRules() map[string]Rule {
	return map[string]Rule{
		"Swamp Palace": swampPalaceCanEnter,
	}
}

func swampPalaceCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Swamp Palace": swampPalaceLocationRules()["Swamp Palace - Boss"],
	}
}

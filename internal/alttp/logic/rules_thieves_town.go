package logic

func thievesTownCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(1)) && items.HasHealth(7) && items.HasABottle())) &&
		items.Has("Moon Pearl") && regions("North West Dark World")
}

func thievesTownLocationRules() map[string]Rule {
	return map[string]Rule{
		"Thieves' Town - Attic": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Key (Thieves Town)") && items.Has("Big Key (Thieves Town)")
		},
		"Thieves' Town - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if settings.LocationHasItem("Thieves' Town - Big Chest", "Key (Thieves Town)") {
				return items.Has("Hammer") && items.Has("Big Key (Thieves Town)")
			}
			return items.Has("Hammer") && items.Has("Key (Thieves Town)") && items.Has("Big Key (Thieves Town)")
		},
		"Thieves' Town - Blind's Cell": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Thieves Town)")
		},
		"Thieves' Town - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return thievesTownCanEnter(items, settings, regions) &&
				items.Has("Key (Thieves Town)") && items.Has("Big Key (Thieves Town)") &&
				CanBeatBlind(items, settings)
		},
		"Thieves' Town - Map Chest":     alwaysAccessible,
		"Thieves' Town - Compass Chest": alwaysAccessible,
		"Thieves' Town - Ambush Chest":  alwaysAccessible,
		"Thieves' Town - Big Key Chest": alwaysAccessible,
	}
}

func thievesTownRegionRules() map[string]Rule {
	return map[string]Rule{
		"Thieves Town": thievesTownCanEnter,
	}
}

func thievesTownCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Thieves Town": thievesTownLocationRules()["Thieves' Town - Boss"],
	}
}

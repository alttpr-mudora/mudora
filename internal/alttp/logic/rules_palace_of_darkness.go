package logic

func palaceOfDarknessCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(1)) && items.HasHealth(7) && items.HasABottle())) &&
		items.Has("Moon Pearl") && regions("North East Dark World")
}

func palaceOfDarknessLocationRules() map[string]Rule {
	hammerAndArrowsAndLamp := func(items *Items, settings *Settings) bool {
		return items.Has("Hammer") && items.CanShootArrows(settings, 1) &&
			items.HasAtLeast("Lamp", settings.LampRequireCount)
	}

	hammerAndArrows := func(items *Items, settings *Settings) bool {
		return items.Has("Hammer") && items.CanShootArrows(settings, 1)
	}

	return map[string]Rule{
		"Palace of Darkness - Shooter Room": alwaysAccessible,
		"Palace of Darkness - The Arena - Ledge": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanShootArrows(settings, 1)
		},
		"Palace of Darkness - Big Key Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if hammerAndArrowsAndLamp(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 6)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 5)
		},
		"Palace of Darkness - The Arena - Bridge": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Key (Palace of Darkness)") || hammerAndArrows(items, settings)
		},
		"Palace of Darkness - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !items.HasAtLeast("Lamp", settings.LampRequireCount) || !items.Has("Big Key (Palace of Darkness)") {
				return false
			}
			if hammerAndArrows(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 6)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 5)
		},
		"Palace of Darkness - Compass Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if hammerAndArrowsAndLamp(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 4)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 3)
		},
		"Palace of Darkness - Harmless Hellway": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if hammerAndArrowsAndLamp(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 6)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 5)
		},
		"Palace of Darkness - Stalfos Basement": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Key (Palace of Darkness)") || hammerAndArrows(items, settings)
		},
		"Palace of Darkness - Dark Basement - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			hasLight := items.HasAtLeast("Lamp", settings.LampRequireCount) ||
				(settings.ItemPlacementAdvanced && items.Has("Fire Rod"))
			if !hasLight {
				return false
			}
			if hammerAndArrowsAndLamp(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 4)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 3)
		},
		"Palace of Darkness - Dark Basement - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			hasLight := items.HasAtLeast("Lamp", settings.LampRequireCount) ||
				(settings.ItemPlacementAdvanced && items.Has("Fire Rod"))
			if !hasLight {
				return false
			}
			if hammerAndArrowsAndLamp(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 4)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 3)
		},
		"Palace of Darkness - Map Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanShootArrows(settings, 1)
		},
		"Palace of Darkness - Dark Maze - Top": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !items.HasAtLeast("Lamp", settings.LampRequireCount) {
				return false
			}
			if hammerAndArrows(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 6)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 5)
		},
		"Palace of Darkness - Dark Maze - Bottom": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !items.HasAtLeast("Lamp", settings.LampRequireCount) {
				return false
			}
			if hammerAndArrows(items, settings) {
				return items.HasAtLeast("Key (Palace of Darkness)", 6)
			}
			return items.HasAtLeast("Key (Palace of Darkness)", 5)
		},
		"Palace of Darkness - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return palaceOfDarknessCanEnter(items, settings, regions) &&
				CanBeatHelmasaurKing(items, settings) &&
				items.Has("Hammer") && items.HasAtLeast("Lamp", settings.LampRequireCount) &&
				items.CanShootArrows(settings, 1) &&
				items.Has("Big Key (Palace of Darkness)") &&
				items.HasAtLeast("Key (Palace of Darkness)", 6)
		},
	}
}

func palaceOfDarknessRegionRules() map[string]Rule {
	return map[string]Rule{
		"Palace of Darkness": palaceOfDarknessCanEnter,
	}
}

func palaceOfDarknessCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Palace of Darkness": palaceOfDarknessLocationRules()["Palace of Darkness - Boss"],
	}
}

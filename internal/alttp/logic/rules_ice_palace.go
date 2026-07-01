package logic

func icePalaceCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(2)) && items.HasHealth(12) &&
				(items.HasBottle(2) || items.HasArmor(1)))) &&
		items.CanMeltThings(settings) &&
		items.Has("Moon Pearl") && items.Has("Flippers") && items.CanLiftDarkRocks()
}

func icePalaceLocationRules() map[string]Rule {
	spikeRoom := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Hookshot")
	}

	hammerLiftAndSpikeRoom := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Hammer") && items.CanLiftRocks() && spikeRoom(items, settings, regions)
	}

	return map[string]Rule{
		"Ice Palace - Spike Room":    spikeRoom,
		"Ice Palace - Big Key Chest": hammerLiftAndSpikeRoom,
		"Ice Palace - Map Chest":     hammerLiftAndSpikeRoom,
		"Ice Palace - Compass Chest": alwaysAccessible,
		"Ice Palace - Iced T Room":   alwaysAccessible,
		"Ice Palace - Freezor Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanMeltThings(settings)
		},
		"Ice Palace - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Ice Palace)")
		},
		"Ice Palace - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return icePalaceCanEnter(items, settings, regions) &&
				items.Has("Hammer") && items.CanLiftRocks() &&
				CanBeatKholdstare(items, settings) &&
				items.Has("Big Key (Ice Palace)") &&
				((settings.ItemPlacementAdvanced && items.Has("Cane of Somaria") && items.Has("Key (Ice Palace)")) ||
					items.HasAtLeast("Key (Ice Palace)", 2))
		},
	}
}

func icePalaceRegionRules() map[string]Rule {
	return map[string]Rule{
		"Ice Palace": icePalaceCanEnter,
	}
}

func icePalaceCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Ice Palace": icePalaceLocationRules()["Ice Palace - Boss"],
	}
}

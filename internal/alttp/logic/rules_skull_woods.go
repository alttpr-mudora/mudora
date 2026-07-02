package logic

func skullWoodsCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(1)) && items.HasHealth(7) && items.HasABottle())) &&
		items.Has("Moon Pearl") && regions("North West Dark World")
}

func skullWoodsLocationRules() map[string]Rule {
	bridgeRoom := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Fire Rod") && items.Has("Moon Pearl")
	}

	return map[string]Rule{
		"Skull Woods - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Skull Woods)")
		},
		"Skull Woods - Bridge Room": bridgeRoom,
		"Skull Woods - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return skullWoodsCanEnter(items, settings, regions) &&
				bridgeRoom(items, settings, regions) &&
				(settings.SwordlessMode || items.HasSword(1)) &&
				items.HasAtLeast("Key (Skull Woods)", 3) &&
				CanBeatMothula(items, settings)
		},
		"Skull Woods - Big Key Chest": alwaysAccessible,
		"Skull Woods - Compass Chest": alwaysAccessible,
		"Skull Woods - Map Chest":     alwaysAccessible,
		"Skull Woods - Pot Prison":    alwaysAccessible,
		"Skull Woods - Pinball Room":  alwaysAccessible,
	}
}

func skullWoodsRegionRules() map[string]Rule {
	return map[string]Rule{
		"Skull Woods": skullWoodsCanEnter,
	}
}

func skullWoodsCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Skull Woods": skullWoodsLocationRules()["Skull Woods - Boss"],
	}
}

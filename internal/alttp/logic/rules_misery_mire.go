package logic

func miseryMireCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(2)) && items.HasHealth(12) &&
				(items.HasBottle(2) || items.HasArmor(1)))) &&
		settings.MiseryMireMedallion != "" && items.Has(settings.MiseryMireMedallion) &&
		(settings.SwordlessMode || items.HasSword(1)) &&
		items.Has("Moon Pearl") &&
		((settings.ItemPlacementAdvanced && items.Has("Pegasus Boots")) || items.Has("Hookshot")) &&
		items.CanKillMostThings(settings, 8) && regions("Mire")
}

func miseryMireLocationRules() map[string]Rule {
	mainLobbyOrBigKey := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Key (Misery Mire)") || items.Has("Big Key (Misery Mire)")
	}

	torchesAndThreeKeys := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.CanLightTorches() && items.HasAtLeast("Key (Misery Mire)", 3)
	}

	return map[string]Rule{
		"Misery Mire - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Misery Mire)")
		},
		"Misery Mire - Spike Chest":   alwaysAccessible,
		"Misery Mire - Main Lobby":    mainLobbyOrBigKey,
		"Misery Mire - Map Chest":     mainLobbyOrBigKey,
		"Misery Mire - Big Key Chest": torchesAndThreeKeys,
		"Misery Mire - Compass Chest": torchesAndThreeKeys,
		"Misery Mire - Bridge Chest":  alwaysAccessible,
		"Misery Mire - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return miseryMireCanEnter(items, settings, regions) &&
				items.Has("Cane of Somaria") && items.HasAtLeast("Lamp", settings.LampRequireCount) &&
				items.Has("Big Key (Misery Mire)") &&
				CanBeatVitreous(items, settings)
		},
	}
}

func miseryMireRegionRules() map[string]Rule {
	return map[string]Rule{
		"Misery Mire": miseryMireCanEnter,
	}
}

func miseryMireCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Misery Mire": miseryMireLocationRules()["Misery Mire - Boss"],
	}
}

package logic

func turtleRockUpper(items *Items, settings *Settings, regions RegionAccess) bool {
	return settings.TurtleRockMedallion != "" && items.Has(settings.TurtleRockMedallion) &&
		(settings.SwordlessMode || items.HasSword(1)) &&
		items.Has("Moon Pearl") && items.Has("Cane of Somaria") &&
		items.Has("Hammer") && items.CanLiftDarkRocks() && regions("East Death Mountain")
}

func turtleRockCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(2)) && items.HasHealth(12) &&
				(items.HasBottle(2) || items.HasArmor(1)))) &&
		turtleRockUpper(items, settings, regions)
}

func turtleRockLocationRules() map[string]Rule {
	eyeBridge := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return turtleRockUpper(items, settings, regions) &&
			items.HasAtLeast("Lamp", settings.LampRequireCount) &&
			items.Has("Cane of Somaria") && items.Has("Big Key (Turtle Rock)") &&
			items.HasAtLeast("Key (Turtle Rock)", 3) &&
			(settings.ItemPlacementAdvanced || items.Has("Cape") || items.Has("Cane of Byrna") || items.CanBlockLasers())
	}

	return map[string]Rule{
		"Turtle Rock - Chain Chomps": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return turtleRockUpper(items, settings, regions) && items.Has("Key (Turtle Rock)")
		},
		"Turtle Rock - Roller Room - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Fire Rod") && items.Has("Cane of Somaria") && turtleRockUpper(items, settings, regions)
		},
		"Turtle Rock - Roller Room - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Fire Rod") && items.Has("Cane of Somaria") && turtleRockUpper(items, settings, regions)
		},
		"Turtle Rock - Compass Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Cane of Somaria") && turtleRockUpper(items, settings, regions)
		},
		"Turtle Rock - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Turtle Rock)") && turtleRockUpper(items, settings, regions) &&
				items.HasAtLeast("Key (Turtle Rock)", 2)
		},
		"Turtle Rock - Big Key Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.HasAtLeast("Key (Turtle Rock)", 2)
		},
		"Turtle Rock - Crystaroller Room": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Turtle Rock)") && turtleRockUpper(items, settings, regions) &&
				items.HasAtLeast("Key (Turtle Rock)", 2)
		},
		"Turtle Rock - Eye Bridge - Bottom Left":  eyeBridge,
		"Turtle Rock - Eye Bridge - Bottom Right": eyeBridge,
		"Turtle Rock - Eye Bridge - Top Left":     eyeBridge,
		"Turtle Rock - Eye Bridge - Top Right":    eyeBridge,
		"Turtle Rock - Boss": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return turtleRockCanEnter(items, settings, regions) &&
				items.HasAtLeast("Key (Turtle Rock)", 4) &&
				items.HasAtLeast("Lamp", settings.LampRequireCount) &&
				items.Has("Big Key (Turtle Rock)") && items.Has("Cane of Somaria") &&
				CanBeatTrinexx(items, settings)
		},
	}
}

func turtleRockRegionRules() map[string]Rule {
	return map[string]Rule{
		"Turtle Rock": turtleRockCanEnter,
	}
}

func turtleRockCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Turtle Rock": turtleRockLocationRules()["Turtle Rock - Boss"],
	}
}

package logic

func deathMountainLocationRules() map[string]Rule {
	return map[string]Rule{
		// West (Light World)
		"Old Man": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.HasAtLeast("Lamp", settings.LampRequireCount)
		},
		"Ether Tablet": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Book of Mudora") &&
				(items.HasSword(2) || (settings.SwordlessMode && items.Has("Hammer"))) &&
				regions("Tower of Hera")
		},
		"Spectacle Rock": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Magic Mirror")
		},
		"Spectacle Rock Cave": alwaysAccessible,

		// East (Light World)
		"Mimic Cave": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hammer") && items.Has("Magic Mirror") &&
				items.HasAtLeast("Key (Turtle Rock)", 2) && regions("Turtle Rock")
		},
		"Floating Island": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Magic Mirror") && items.Has("Moon Pearl") &&
				items.CanBombThings() && items.CanLiftRocks() &&
				regions("East Dark World Death Mountain")
		},
		"Spiral Cave":                    alwaysAccessible,
		"Paradox Cave Lower - Far Left":  alwaysAccessible,
		"Paradox Cave Lower - Left":      alwaysAccessible,
		"Paradox Cave Lower - Right":     alwaysAccessible,
		"Paradox Cave Lower - Far Right": alwaysAccessible,
		"Paradox Cave Lower - Middle":    alwaysAccessible,
		"Paradox Cave Upper - Left":      alwaysAccessible,
		"Paradox Cave Upper - Right":     alwaysAccessible,

		// West (Dark World)
		"Spike Cave": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl") && items.Has("Hammer") && items.CanLiftRocks() &&
				((items.CanExtendMagic(settings, 2.0) && items.Has("Cape")) || items.Has("Cane of Byrna"))
		},

		// East (Dark World)
		"Superbunny Cave - Top": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Superbunny Cave - Bottom": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Hookshot Cave - Top Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && items.Has("Moon Pearl") && items.CanLiftRocks()
		},
		"Hookshot Cave - Top Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && items.Has("Moon Pearl") && items.CanLiftRocks()
		},
		"Hookshot Cave - Bottom Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && items.Has("Moon Pearl") && items.CanLiftRocks()
		},
		"Hookshot Cave - Bottom Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return (items.Has("Hookshot") || (settings.ItemPlacementAdvanced && items.Has("Pegasus Boots"))) &&
				items.Has("Moon Pearl") && items.CanLiftRocks()
		},
	}
}

func deathMountainRegionRules() map[string]Rule {
	westDeathMountain := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") &&
			(items.CanFly() || (items.CanLiftRocks() && items.HasAtLeast("Lamp", settings.LampRequireCount)))
	}

	eastDeathMountain := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") &&
			((items.Has("Hookshot") && regions("West Death Mountain")) ||
				(items.Has("Hammer") && regions("Tower of Hera")))
	}

	westDarkWorldDeathMountain := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") && regions("West Death Mountain")
	}

	eastDarkWorldDeathMountain := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") && items.CanLiftDarkRocks() && regions("East Death Mountain")
	}

	return map[string]Rule{
		"West Death Mountain":            westDeathMountain,
		"East Death Mountain":            eastDeathMountain,
		"West Dark World Death Mountain": westDarkWorldDeathMountain,
		"East Dark World Death Mountain": eastDarkWorldDeathMountain,
	}
}

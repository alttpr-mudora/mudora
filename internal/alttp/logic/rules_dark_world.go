package logic

func darkWorldLocationRules() map[string]Rule {
	blacksmith := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return (settings.ItemPlacementAdvanced || items.Has("Magic Mirror")) &&
			items.Has("Moon Pearl") && items.CanLiftDarkRocks()
	}

	rules := map[string]Rule{
		// North West
		"Blacksmith": blacksmith,
		"Purple Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return blacksmith(items, settings, regions) && items.Has("Moon Pearl") && items.CanLiftDarkRocks()
		},
		"Brewery": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.Has("Moon Pearl")
		},
		"C-Shaped House": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Chest Game": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Hammer Pegs": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hammer") && items.Has("Moon Pearl") && items.CanLiftDarkRocks()
		},
		"Bumper Cave": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl") &&
				(settings.ItemPlacementAdvanced || items.Has("Hookshot")) &&
				items.CanLiftRocks() && items.Has("Cape")
		},

		// North East
		"Catfish": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl") && items.CanLiftRocks()
		},
		"Pyramid": alwaysAccessible,
		"Pyramid Fairy - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Crystal 5") && items.Has("Crystal 6") && regions("South Dark World") &&
				items.Has("Moon Pearl") && items.Has("Hammer")
		},
		"Pyramid Fairy - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Crystal 5") && items.Has("Crystal 6") && regions("South Dark World") &&
				items.Has("Moon Pearl") && items.Has("Hammer")
		},

		// South
		"Hype Cave - Top": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Hype Cave - Middle Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Hype Cave - Middle Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Hype Cave - Bottom": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Hype Cave - NPC": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Stumpy": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl") || items.Has("Magic Mirror")
		},
		"Digging Game": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
	}

	return rules
}

func darkWorldRegionRules() map[string]Rule {
	northWestDarkWorld := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") && items.Has("Moon Pearl") &&
			((regions("North East Dark World") && items.Has("Hookshot") &&
				(items.CanLiftRocks() || items.Has("Hammer") || items.Has("Flippers"))) ||
				(items.Has("Hammer") && items.CanLiftRocks()) ||
				items.CanLiftDarkRocks())
	}

	northEastDarkWorld := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") &&
			(items.Has("Defeat Agahnim") ||
				(items.Has("Hammer") && items.CanLiftRocks() && items.Has("Moon Pearl")) ||
				(items.CanLiftDarkRocks() && items.Has("Moon Pearl") &&
					(items.Has("Hammer") || items.Has("Flippers"))))
	}

	southDarkWorld := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda") &&
			((items.Has("Moon Pearl") && regions("North East Dark World") && items.Has("Hammer")) ||
				regions("North West Dark World"))
	}

	return map[string]Rule{
		"North West Dark World": northWestDarkWorld,
		"North East Dark World": northEastDarkWorld,
		"South Dark World":      southDarkWorld,
	}
}

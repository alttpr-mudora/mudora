package logic

func lightWorldLocationRules() map[string]Rule {
	rules := map[string]Rule{
		// North West
		"Master Sword Pedestal": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return (settings.ItemPlacementAdvanced || items.Has("Book of Mudora")) &&
				items.Has("Pendant of Power") && items.Has("Pendant of Wisdom") && items.Has("Pendant of Courage")
		},
		"King's Tomb": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Pegasus Boots") &&
				(items.CanLiftDarkRocks() ||
					(items.Has("Magic Mirror") && regions("North West Dark World") && items.Has("Moon Pearl")))
		},
		"Pegasus Rocks": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Pegasus Boots")
		},
		"Magic Bat": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Powder") &&
				(items.Has("Hammer") ||
					(items.Has("Magic Mirror") && items.Has("Moon Pearl") &&
						items.CanLiftDarkRocks() && regions("North West Dark World")))
		},
		"Sick Kid": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.HasABottle()
		},
		"Lumberjack Tree": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Defeat Agahnim") && items.Has("Pegasus Boots")
		},
		"Graveyard Ledge": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Magic Mirror") && regions("North West Dark World") && items.Has("Moon Pearl")
		},
		"Kakariko Tavern":             alwaysAccessible,
		"Chicken House":               alwaysAccessible,
		"Kakariko Well - Top":         alwaysAccessible,
		"Kakariko Well - Left":        alwaysAccessible,
		"Kakariko Well - Middle":      alwaysAccessible,
		"Kakariko Well - Right":       alwaysAccessible,
		"Kakariko Well - Bottom":      alwaysAccessible,
		"Blind's Hideout - Top":       alwaysAccessible,
		"Blind's Hideout - Left":      alwaysAccessible,
		"Blind's Hideout - Right":     alwaysAccessible,
		"Blind's Hideout - Far Left":  alwaysAccessible,
		"Blind's Hideout - Far Right": alwaysAccessible,
		"Bottle Merchant":             alwaysAccessible,
		"Lost Woods Hideout":          alwaysAccessible,
		"Mushroom":                    alwaysAccessible,

		// North East
		"Sahasrahla": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Pendant of Courage")
		},
		"King Zora": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanLiftRocks() || items.Has("Flippers")
		},
		"Potion Shop": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Mushroom")
		},
		"Zora's Ledge": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Flippers")
		},
		"Waterfall Fairy - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Flippers")
		},
		"Waterfall Fairy - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Flippers")
		},
		"Sahasrahla's Hut - Left":   alwaysAccessible,
		"Sahasrahla's Hut - Middle": alwaysAccessible,
		"Sahasrahla's Hut - Right":  alwaysAccessible,

		// South
		"Mini Moldorm Cave - Far Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.CanKillMostThings(settings, 5)
		},
		"Mini Moldorm Cave - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.CanKillMostThings(settings, 5)
		},
		"Mini Moldorm Cave - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.CanKillMostThings(settings, 5)
		},
		"Mini Moldorm Cave - Far Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.CanKillMostThings(settings, 5)
		},
		"Mini Moldorm Cave - NPC": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings() && items.CanKillMostThings(settings, 5)
		},
		"Hobo": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Flippers")
		},
		"Bombos Tablet": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Book of Mudora") &&
				(items.HasSword(2) || (settings.SwordlessMode && items.Has("Hammer"))) &&
				items.Has("Magic Mirror") && regions("South Dark World")
		},
		"Cave 45": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Magic Mirror") && regions("South Dark World")
		},
		"Checkerboard Cave": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanLiftRocks() && items.Has("Magic Mirror") && regions("Mire")
		},
		"Library": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Pegasus Boots")
		},
		"Desert Ledge": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return regions("Desert Palace")
		},
		"Lake Hylia Island": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Flippers") && items.Has("Magic Mirror") && items.Has("Moon Pearl") &&
				regions("North East Dark World")
		},
		"Flute Spot": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Shovel")
		},
		"Aginah's Cave": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.CanBombThings()
		},
		"Floodgate Chest": alwaysAccessible,
		"Link's House":    alwaysAccessible,
		"Ice Rod Cave":    alwaysAccessible,
		"Maze Race":       alwaysAccessible,
		"Sunken Treasure": alwaysAccessible,
	}

	return rules
}

func lightWorldRegionRules() map[string]Rule {
	rescuedZelda := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Rescue Zelda")
	}

	return map[string]Rule{
		"North West Light World": rescuedZelda,
		"North East Light World": rescuedZelda,
		"South Light World":      rescuedZelda,
	}
}

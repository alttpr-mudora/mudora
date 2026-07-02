package logic

func ganonsTowerCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.Has("Rescue Zelda") &&
		(settings.ItemPlacementAdvanced ||
			((settings.SwordlessMode || items.HasSword(2)) && items.HasHealth(12) &&
				(items.HasBottle(2) || items.HasArmor(1)))) &&
		items.Has("Moon Pearl") && items.CrystalCount() >= settings.CrystalsRequiredForTower &&
		regions("East Dark World Death Mountain")
}

func ganonsTowerLocationRules() map[string]Rule {
	hammerAndHookshot := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.Has("Hammer") && items.Has("Hookshot")
	}

	randomizerRoom := func(siblings ...string) Rule {
		return func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !hammerAndHookshot(items, settings, regions) {
				return false
			}
			for _, s := range siblings {
				if settings.LocationHasItem(s, "Big Key (Ganon's Tower)") {
					return items.HasAtLeast("Key (Ganon's Tower)", 3)
				}
			}
			return items.HasAtLeast("Key (Ganon's Tower)", 4)
		}
	}

	hammerHookshotOrFireRodSomaria := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return hammerAndHookshot(items, settings, regions) || (items.Has("Fire Rod") && items.Has("Cane of Somaria"))
	}

	compassRoom := func(siblings ...string) Rule {
		return func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !items.Has("Fire Rod") || !items.Has("Cane of Somaria") {
				return false
			}
			for _, s := range siblings {
				if settings.LocationHasItem(s, "Big Key (Ganon's Tower)") {
					return items.HasAtLeast("Key (Ganon's Tower)", 3)
				}
			}
			return items.HasAtLeast("Key (Ganon's Tower)", 4)
		}
	}

	bigKeyRoom := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return hammerHookshotOrFireRodSomaria(items, settings, regions) &&
			items.HasAtLeast("Key (Ganon's Tower)", 3) && CanBeatArmosKnights(items, settings)
	}

	preMoldormChest := func(items *Items, settings *Settings, regions RegionAccess) bool {
		return items.CanShootArrows(settings, 1) && items.CanLightTorches() &&
			items.Has("Big Key (Ganon's Tower)") && items.HasAtLeast("Key (Ganon's Tower)", 3) &&
			CanBeatLanmolas(items, settings)
	}

	return map[string]Rule{
		"Ganon's Tower - Bob's Torch": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Pegasus Boots")
		},
		"Ganon's Tower - DMs Room - Top Left":     hammerAndHookshot,
		"Ganon's Tower - DMs Room - Top Right":    hammerAndHookshot,
		"Ganon's Tower - DMs Room - Bottom Left":  hammerAndHookshot,
		"Ganon's Tower - DMs Room - Bottom Right": hammerAndHookshot,

		"Ganon's Tower - Randomizer Room - Top Left": randomizerRoom(
			"Ganon's Tower - Randomizer Room - Top Right",
			"Ganon's Tower - Randomizer Room - Bottom Left",
			"Ganon's Tower - Randomizer Room - Bottom Right",
		),
		"Ganon's Tower - Randomizer Room - Top Right": randomizerRoom(
			"Ganon's Tower - Randomizer Room - Top Left",
			"Ganon's Tower - Randomizer Room - Bottom Left",
			"Ganon's Tower - Randomizer Room - Bottom Right",
		),
		"Ganon's Tower - Randomizer Room - Bottom Left": randomizerRoom(
			"Ganon's Tower - Randomizer Room - Top Right",
			"Ganon's Tower - Randomizer Room - Top Left",
			"Ganon's Tower - Randomizer Room - Bottom Right",
		),
		"Ganon's Tower - Randomizer Room - Bottom Right": randomizerRoom(
			"Ganon's Tower - Randomizer Room - Top Right",
			"Ganon's Tower - Randomizer Room - Top Left",
			"Ganon's Tower - Randomizer Room - Bottom Left",
		),

		"Ganon's Tower - Firesnake Room": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !hammerAndHookshot(items, settings, regions) {
				return false
			}
			randomizerRoomsHaveBigKey := settings.LocationHasItem("Ganon's Tower - Randomizer Room - Top Right", "Big Key (Ganon's Tower)") ||
				settings.LocationHasItem("Ganon's Tower - Randomizer Room - Top Left", "Big Key (Ganon's Tower)") ||
				settings.LocationHasItem("Ganon's Tower - Randomizer Room - Bottom Left", "Big Key (Ganon's Tower)") ||
				settings.LocationHasItem("Ganon's Tower - Randomizer Room - Bottom Right", "Big Key (Ganon's Tower)")
			firesnakeSelfLocked := settings.LocationHasItem("Ganon's Tower - Firesnake Room", "Key (Ganon's Tower)")
			if (randomizerRoomsHaveBigKey || firesnakeSelfLocked) && items.HasAtLeast("Key (Ganon's Tower)", 2) {
				return true
			}
			return items.HasAtLeast("Key (Ganon's Tower)", 3)
		},
		"Ganon's Tower - Map Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			if !items.Has("Hammer") || !(items.Has("Hookshot") || (settings.ItemPlacementAdvanced && items.Has("Pegasus Boots"))) {
				return false
			}
			selfLocked := settings.LocationHasItem("Ganon's Tower - Map Chest", "Big Key (Ganon's Tower)") ||
				settings.LocationHasItem("Ganon's Tower - Map Chest", "Key (Ganon's Tower)")
			if selfLocked {
				return items.HasAtLeast("Key (Ganon's Tower)", 3)
			}
			return items.HasAtLeast("Key (Ganon's Tower)", 4)
		},
		"Ganon's Tower - Big Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Big Key (Ganon's Tower)") && items.HasAtLeast("Key (Ganon's Tower)", 3) &&
				hammerHookshotOrFireRodSomaria(items, settings, regions)
		},
		"Ganon's Tower - Bob's Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return hammerHookshotOrFireRodSomaria(items, settings, regions) &&
				items.HasAtLeast("Key (Ganon's Tower)", 3) &&
				(settings.ItemPlacementAdvanced || items.Has("Fire Rod") || (items.Has("Ether") && items.HasSword(1)))
		},
		"Ganon's Tower - Tile Room": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Cane of Somaria")
		},
		"Ganon's Tower - Hope Room - Left":  alwaysAccessible,
		"Ganon's Tower - Hope Room - Right": alwaysAccessible,

		"Ganon's Tower - Compass Room - Top Left": compassRoom(
			"Ganon's Tower - Compass Room - Top Right",
			"Ganon's Tower - Compass Room - Bottom Left",
			"Ganon's Tower - Compass Room - Bottom Right",
		),
		"Ganon's Tower - Compass Room - Top Right": compassRoom(
			"Ganon's Tower - Compass Room - Top Left",
			"Ganon's Tower - Compass Room - Bottom Left",
			"Ganon's Tower - Compass Room - Bottom Right",
		),
		"Ganon's Tower - Compass Room - Bottom Left": compassRoom(
			"Ganon's Tower - Compass Room - Top Right",
			"Ganon's Tower - Compass Room - Top Left",
			"Ganon's Tower - Compass Room - Bottom Right",
		),
		"Ganon's Tower - Compass Room - Bottom Right": compassRoom(
			"Ganon's Tower - Compass Room - Top Right",
			"Ganon's Tower - Compass Room - Top Left",
			"Ganon's Tower - Compass Room - Bottom Left",
		),

		"Ganon's Tower - Big Key Chest":               bigKeyRoom,
		"Ganon's Tower - Big Key Room - Left":         bigKeyRoom,
		"Ganon's Tower - Big Key Room - Right":        bigKeyRoom,
		"Ganon's Tower - Mini Helmasaur Room - Left":  preMoldormChest,
		"Ganon's Tower - Mini Helmasaur Room - Right": preMoldormChest,
		"Ganon's Tower - Pre-Moldorm Chest":           preMoldormChest,
		"Ganon's Tower - Moldorm Chest": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Hookshot") && items.CanShootArrows(settings, 1) && items.CanLightTorches() &&
				items.Has("Big Key (Ganon's Tower)") && items.HasAtLeast("Key (Ganon's Tower)", 4) &&
				CanBeatLanmolas(items, settings) && CanBeatMoldorm(items, settings)
		},
	}
}

func ganonsTowerRegionRules() map[string]Rule {
	return map[string]Rule{
		"Ganons Tower": ganonsTowerCanEnter,
	}
}

func ganonsTowerCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Ganons Tower": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return ganonsTowerCanEnter(items, settings, regions) &&
				ganonsTowerLocationRules()["Ganon's Tower - Moldorm Chest"](items, settings, regions) &&
				CanBeatAgahnim2(items, settings)
		},
	}
}

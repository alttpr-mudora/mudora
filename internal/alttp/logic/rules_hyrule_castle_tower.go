package logic

func hyruleCastleTowerLocationRules() map[string]Rule {
	return map[string]Rule{
		"Castle Tower - Dark Maze": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.HasAtLeast("Lamp", settings.LampRequireCount) && items.Has("Key (Castle Tower)")
		},
		"Castle Tower - Room 03": alwaysAccessible,
	}
}

func hyruleCastleTowerCanEnter(items *Items, settings *Settings, regions RegionAccess) bool {
	return items.CanKillMostThings(settings, 8) &&
		items.Has("Rescue Zelda") &&
		(items.Has("Cape") || items.HasSword(2) || (settings.SwordlessMode && items.Has("Hammer")))
}

func hyruleCastleTowerRegionRules() map[string]Rule {
	return map[string]Rule{
		"Castle Tower": hyruleCastleTowerCanEnter,
	}
}

func hyruleCastleTowerCompletionRules() map[string]Rule {
	return map[string]Rule{
		"Castle Tower": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return hyruleCastleTowerCanEnter(items, settings, regions) &&
				items.HasAtLeast("Key (Castle Tower)", 2) &&
				items.HasAtLeast("Lamp", settings.LampRequireCount) &&
				(items.HasSword(1) || (settings.SwordlessMode && (items.Has("Hammer") || items.Has("Bug Net"))))
		},
	}
}

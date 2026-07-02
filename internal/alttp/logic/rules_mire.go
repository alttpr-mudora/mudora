package logic

func mireLocationRules() map[string]Rule {
	return map[string]Rule{
		"Mire Shed - Left": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
		"Mire Shed - Right": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Moon Pearl")
		},
	}
}

func mireRegionRules() map[string]Rule {
	return map[string]Rule{
		"Mire": func(items *Items, settings *Settings, regions RegionAccess) bool {
			return items.Has("Rescue Zelda") && items.CanLiftDarkRocks() && items.CanFly()
		},
	}
}

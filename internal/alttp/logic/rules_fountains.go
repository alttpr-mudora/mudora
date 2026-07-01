package logic

func fountainsLocationRules() map[string]Rule {
	return map[string]Rule{
		"Waterfall Bottle": alwaysAccessible,
		"Pyramid Bottle":   alwaysAccessible,
	}
}

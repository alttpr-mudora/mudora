package logic

func CanBeatArmosKnights(items *Items, settings *Settings) bool {
	return items.HasSword(1) || items.Has("Hammer") || items.CanShootArrows(settings, 1) ||
		items.Has("Boomerang") || items.Has("Boomerang (Red)") ||
		(items.CanExtendMagic(settings, 4) && (items.Has("Fire Rod") || items.Has("Ice Rod"))) ||
		(items.CanExtendMagic(settings, 2) && (items.Has("Cane of Byrna") || items.Has("Cane of Somaria")))
}

func CanBeatLanmolas(items *Items, settings *Settings) bool {
	return items.HasSword(1) || items.Has("Hammer") ||
		items.CanShootArrows(settings, 1) || items.Has("Fire Rod") || items.Has("Ice Rod") ||
		items.Has("Cane of Byrna") || items.Has("Cane of Somaria")
}

func CanBeatMoldorm(items *Items, settings *Settings) bool {
	return items.HasSword(1) || items.Has("Hammer")
}

func CanBeatAgahnim(items *Items, settings *Settings) bool {
	return items.HasSword(1) || items.Has("Hammer") || items.Has("Bug Net")
}

func CanBeatHelmasaurKing(items *Items, settings *Settings) bool {
	return (items.CanBombThings() || items.Has("Hammer")) &&
		(items.HasSword(2) || items.CanShootArrows(settings, 1) ||
			(settings.ItemPlacementAdvanced && items.HasSword(1)))
}

func CanBeatArrghus(items *Items, settings *Settings) bool {
	return (settings.ItemPlacementAdvanced || settings.SwordlessMode || items.HasSword(2)) &&
		items.Has("Hookshot") && (items.Has("Hammer") || items.HasSword(1) ||
		((items.CanExtendMagic(settings, 2) || items.CanShootArrows(settings, 1)) &&
			(items.Has("Fire Rod") || items.Has("Ice Rod"))))
}

func CanBeatMothula(items *Items, settings *Settings) bool {
	return (settings.ItemPlacementAdvanced || items.HasSword(2) ||
		(items.CanExtendMagic(settings, 2) && items.Has("Fire Rod"))) &&
		(items.HasSword(1) || items.Has("Hammer") ||
			(items.CanExtendMagic(settings, 2) && (items.Has("Fire Rod") || items.Has("Cane of Somaria") ||
				items.Has("Cane of Byrna"))) ||
			items.CanGetGoodBee())
}

func CanBeatBlind(items *Items, settings *Settings) bool {
	return (settings.ItemPlacementAdvanced || settings.SwordlessMode ||
		(items.HasSword(1) && (items.Has("Cape") || items.Has("Cane of Byrna")))) &&
		(items.HasSword(1) || items.Has("Hammer") || items.Has("Cane of Somaria") || items.Has("Cane of Byrna"))
}

func CanBeatKholdstare(items *Items, settings *Settings) bool {
	return (settings.ItemPlacementAdvanced || items.HasSword(2) ||
		(items.CanExtendMagic(settings, 3) && items.Has("Fire Rod")) ||
		(items.Has("Bombos") && (settings.SwordlessMode || items.HasSword(1)) &&
			items.CanExtendMagic(settings, 2) && items.Has("Fire Rod"))) &&
		items.CanMeltThings(settings) && (items.Has("Hammer") || items.HasSword(1) ||
		(items.CanExtendMagic(settings, 3) && items.Has("Fire Rod")) ||
		(items.CanExtendMagic(settings, 2) && items.Has("Fire Rod") && items.Has("Bombos") && settings.SwordlessMode))
}

func CanBeatVitreous(items *Items, settings *Settings) bool {
	return (settings.ItemPlacementAdvanced || items.HasSword(2) || items.CanShootArrows(settings, 1)) &&
		(items.Has("Hammer") || items.HasSword(1) || items.CanShootArrows(settings, 1))
}

func CanBeatTrinexx(items *Items, settings *Settings) bool {
	return items.Has("Fire Rod") && items.Has("Ice Rod") &&
		(settings.ItemPlacementAdvanced || settings.SwordlessMode || items.HasSword(3) ||
			(items.CanExtendMagic(settings, 2) && items.HasSword(2))) &&
		(items.HasSword(3) || items.Has("Hammer") ||
			(items.CanExtendMagic(settings, 2) && items.HasSword(2)) ||
			(items.CanExtendMagic(settings, 4) && items.HasSword(1)))
}

func CanBeatAgahnim2(items *Items, settings *Settings) bool {
	return items.HasSword(1) || items.Has("Hammer") || items.Has("Bug Net")
}

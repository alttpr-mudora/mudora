package logic

var bottleNames = []string{
	"Bottle",
	"Bottle (Red Potion)",
	"Bottle (Green Potion)",
	"Bottle (Blue Potion)",
	"Bottle (Bee)",
	"Bottle (Super bee)",
	"Bottle (Faerie)",
}

func (i *Items) HasHealth(minimum float64) bool {
	return float64(i.counts["Heart Container"])+float64(i.counts["Piece of Heart"])*0.25 >= minimum
}

func (i *Items) BottleCount() int {
	count := 0
	for _, name := range bottleNames {
		count += i.counts[name]
	}

	return count
}

func (i *Items) HasBottle(atLeast int) bool {
	return i.BottleCount() >= atLeast
}

func (i *Items) HasABottle() bool {
	return i.BottleCount() >= 1
}

func (i *Items) HasSword(minLevel int) bool {
	switch {
	case minLevel >= 4:
		return i.HasAtLeast("Progressive Sword", 4)
	case minLevel == 3:
		return i.HasAtLeast("Progressive Sword", 3)
	case minLevel == 2:
		return i.HasAtLeast("Progressive Sword", 2) || i.Has("Master Sword")
	default:
		return i.HasAtLeast("Progressive Sword", 1) || i.Has("Master Sword")
	}
}

func (i *Items) HasArmor(minLevel int) bool {
	if minLevel >= 2 {
		return i.HasAtLeast("Progressive Armor", 2)
	}

	return i.HasAtLeast("Progressive Armor", 1)
}

func (i *Items) CanLiftRocks() bool {
	return i.HasAtLeast("Progressive Glove", 1)
}

func (i *Items) CanLiftDarkRocks() bool {
	return i.HasAtLeast("Progressive Glove", 2)
}

func (i *Items) CanLightTorches() bool {
	return i.Has("Fire Rod") || i.Has("Lamp")
}

func (i *Items) CanMeltThings(s *Settings) bool {
	return i.Has("Fire Rod") || (i.Has("Bombos") && (s.SwordlessMode || i.HasSword(1)))
}

// CanFly assumes Standard/Open mode; PHP's Inverted-only canActivateOcarina
// gate always passes outside Inverted, so it's omitted here.
func (i *Items) CanFly() bool {
	return i.Has("Ocarina") || i.Has("Flute")
}

func (i *Items) CanSpinSpeed() bool {
	return i.Has("Pegasus Boots") && (i.HasSword(1) || i.Has("Hookshot"))
}

func (i *Items) CanAcquireFairy(s *Settings) bool {
	return s.CatchableFairies
}

func (i *Items) CanBunnyRevive(s *Settings) bool {
	return i.HasABottle() && i.Has("Bug Net") && i.CanAcquireFairy(s)
}

func (i *Items) CanShootArrows(s *Settings, minLevel int) bool {
	if minLevel >= 2 {
		return i.HasAtLeast("Progressive Bow", 2) ||
			(i.Has("Silver Arrow Upgrade") && (i.Has("Bow") || i.Has("Progressive Bow")))
	}

	return i.Has("Bow") || i.Has("Progressive Bow")
}

func (i *Items) CanBlockLasers() bool {
	return i.HasAtLeast("Progressive Shield", 3)
}

func (i *Items) CanExtendMagic(s *Settings, bars float64) bool {
	baseMagic := 1.0
	switch {
	case i.Has("Quarter Magic"):
		baseMagic = 4.0
	case i.Has("Half Magic"):
		baseMagic = 2.0
	}

	bottleMagic := baseMagic * float64(i.BottleCount()) * s.BottleFillMagicRatio

	return baseMagic+bottleMagic >= bars
}

func (i *Items) GlitchedLinkInDarkWorld() bool {
	return i.Has("Moon Pearl") || i.HasABottle()
}

func (i *Items) CanBombThings() bool {
	return true
}

func (i *Items) CanGetGoodBee() bool {
	return i.Has("Bug Net") && i.HasABottle() && (i.Has("Pegasus Boots") || (i.HasSword(1) && i.Has("Quake")))
}

func (i *Items) CanKillEscapeThings(s *Settings) bool {
	return i.HasSword(1) ||
		i.Has("Cane of Somaria") ||
		(i.Has("Bombs (10)") && s.EnemizerDefaultHealth) ||
		(i.Has("Cane of Byrna") && s.EnemizerDefaultHealth) ||
		i.CanShootArrows(s, 1) ||
		i.Has("Hammer") ||
		i.Has("Fire Rod") ||
		s.IgnoreCanKillEscapeThings
}

func (i *Items) CanKillMostThings(s *Settings, enemies int) bool {
	return i.HasSword(1) ||
		i.Has("Cane of Somaria") ||
		(i.CanBombThings() && enemies < 6 && s.EnemizerDefaultHealth) ||
		(i.Has("Cane of Byrna") && (enemies < 6 || i.CanExtendMagic(s, 2.0)) && s.EnemizerDefaultHealth) ||
		i.CanShootArrows(s, 1) ||
		i.Has("Hammer") ||
		i.Has("Fire Rod")
}

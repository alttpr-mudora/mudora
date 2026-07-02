package logic

type Settings struct {
	IgnoreCanKillEscapeThings bool
	EnemizerDefaultHealth     bool
	SwordlessMode             bool
	CatchableFairies          bool
	RupeeBow                  bool
	BottleFillMagicRatio      float64
	ItemPlacementAdvanced     bool
	LampRequireCount          int
	CanBootsClip              bool
	CanSuperSpeed             bool
	CanOneFrameClipOW         bool

	// Per-seed required medallion; empty means unknown/unsatisfied.
	MiseryMireMedallion string
	TurtleRockMedallion string

	CrystalsRequiredForTower int

	// Per-seed location -> item, from rom.Inspect. Nil means unknown.
	LocationItems map[string]string
}

func (s *Settings) LocationHasItem(location, item string) bool {
	return s.LocationItems[location] == item
}

func DefaultSettings() *Settings {
	return &Settings{
		IgnoreCanKillEscapeThings: false,
		EnemizerDefaultHealth:     true,
		SwordlessMode:             false,
		CatchableFairies:          true,
		RupeeBow:                  false,
		BottleFillMagicRatio:      1.0,
		ItemPlacementAdvanced:     true,
		LampRequireCount:          1,
		CanBootsClip:              false,
		CanSuperSpeed:             false,
		CanOneFrameClipOW:         false,
		CrystalsRequiredForTower:  7,
	}
}

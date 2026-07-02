package logic

import "testing"

func TestHyruleCastleEscapeGranularity(t *testing.T) {
	rules := hyruleCastleEscapeLocationRules()
	settings := DefaultSettings()

	weaponOnly := NewItems()
	weaponOnly.Add("Hammer")

	weaponAndKey := NewItems()
	weaponAndKey.Add("Hammer")
	weaponAndKey.Add("Key (Hyrule Castle)")

	weaponAndKeyAndLamp := NewItems()
	weaponAndKeyAndLamp.Add("Hammer")
	weaponAndKeyAndLamp.Add("Key (Hyrule Castle)")
	weaponAndKeyAndLamp.Add("Lamp")

	gloveOnly := NewItems()
	gloveOnly.Add("Progressive Glove")

	cases := []struct {
		location string
		items    *Items
		want     bool
	}{
		{"Sanctuary", NewItems(), true},
		{"Hyrule Castle - Map Chest", NewItems(), true},
		{"Secret Passage", NewItems(), true},
		{"Link's Uncle", NewItems(), true},

		{"Hyrule Castle - Boomerang Chest", NewItems(), false},
		{"Hyrule Castle - Boomerang Chest", weaponOnly, false},
		{"Hyrule Castle - Boomerang Chest", weaponAndKey, true},
		{"Hyrule Castle - Zelda's Cell", weaponAndKey, true},

		{"Sewers - Dark Cross", NewItems(), false},
		{"Sewers - Dark Cross", weaponOnly, false},
		{"Sewers - Dark Cross", weaponAndKeyAndLamp, true},

		{"Sewers - Secret Room - Left", weaponAndKey, false},
		{"Sewers - Secret Room - Left", weaponAndKeyAndLamp, true},
		{"Sewers - Secret Room - Left", gloveOnly, true},
		{"Sewers - Secret Room - Middle", gloveOnly, true},
		{"Sewers - Secret Room - Right", gloveOnly, true},
	}

	for _, c := range cases {
		rule, ok := rules[c.location]
		if !ok {
			t.Fatalf("no rule registered for %q", c.location)
		}

		if got := rule(c.items, settings, nil); got != c.want {
			t.Errorf("%q: got %v, want %v", c.location, got, c.want)
		}
	}
}

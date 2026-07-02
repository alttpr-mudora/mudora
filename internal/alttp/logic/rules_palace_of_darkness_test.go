package logic

import "testing"

func TestPalaceOfDarknessBigKeyChestSelfLock(t *testing.T) {
	rules := palaceOfDarknessLocationRules()
	rule := rules["Palace of Darkness - Big Key Chest"]

	oneKey := NewItems()
	oneKey.Add("Key (Palace of Darkness)")

	fourKeys := NewItems()
	for i := 0; i < 4; i++ {
		fourKeys.Add("Key (Palace of Darkness)")
	}

	unknown := DefaultSettings()
	if got := rule(oneKey, unknown, nil); got != false {
		t.Errorf("1 key, no self-lock data: got %v, want false", got)
	}

	selfLocked := DefaultSettings()
	selfLocked.LocationItems = map[string]string{
		"Palace of Darkness - Big Key Chest": "Key (Palace of Darkness)",
	}
	if got := rule(oneKey, selfLocked, nil); got != true {
		t.Errorf("1 key, self-locked: got %v, want true", got)
	}

	notSelfLocked := DefaultSettings()
	notSelfLocked.LocationItems = map[string]string{
		"Palace of Darkness - Big Key Chest": "Big Key (Palace of Darkness)",
	}
	if got := rule(fourKeys, notSelfLocked, nil); got != false {
		t.Errorf("4 keys, not self-locked (needs 5): got %v, want false", got)
	}
}

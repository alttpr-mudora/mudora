package graph

import (
	"testing"

	"github.com/alttpr-mudora/mudora/internal/alttp/logic"
)

func TestBuildStartsWithRescueZeldaPreCollected(t *testing.T) {
	g := Build(map[string]string{}, logic.DefaultSettings())

	items := itemsFor(g.Start, g.ItemAt)
	if !items.Has("Rescue Zelda") {
		t.Fatal("expected Rescue Zelda to be pre-collected at the start state in Open mode")
	}
}

func TestBuildStatesStrictlyGrow(t *testing.T) {
	itemAt := map[string]string{
		"Hyrule Castle - Map Chest": "Key (Hyrule Castle)",
		"Secret Passage":            "Powder",
		"Link's Uncle":              "Flute",
	}

	g := Build(itemAt, logic.DefaultSettings())

	seenSizes := make(map[int]bool)
	for i, s := range g.Order {
		n := len(s.locations())
		if i > 0 && n <= len(g.Order[i-1].locations()) {
			t.Fatalf("state %d did not strictly grow: %d -> %d", i, len(g.Order[i-1].locations()), n)
		}
		if seenSizes[n] {
			t.Fatalf("duplicate state size %d in chain", n)
		}
		seenSizes[n] = true
	}

	if len(g.Order) != len(g.Edges)+1 {
		t.Fatalf("expected a single chain (states = edges+1), got %d states and %d edges", len(g.Order), len(g.Edges))
	}
}

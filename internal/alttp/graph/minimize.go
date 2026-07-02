package graph

import (
	"sort"

	"github.com/alttpr-mudora/mudora/internal/alttp/logic"
)

// MinimalPathTo builds the shortest-in-checks path to goal (e.g. "Ganons
// Tower"). It runs a full completion pass to confirm the goal is reachable
// at all, then greedily drops every collected location that turns out to
// be unnecessary: removing it and replaying the remaining set through
// Build still reaches the goal. That replay (not a flat sum of the
// remaining items) is what necessity is checked against, since a location
// can be a real prerequisite for reaching another one even if it's never
// directly referenced by the goal's own rule.
func MinimalPathTo(itemAt map[string]string, settings *logic.Settings, goal string) *Graph {
	itemAt = withVirtualItems(itemAt)
	full := Build(itemAt, settings)

	locs := full.Order[len(full.Order)-1].locations()
	sort.Strings(locs)

	necessary := make(map[string]bool, len(locs))
	for _, loc := range locs {
		if loc != goal {
			necessary[loc] = true
		}
	}

	reaches := func() bool {
		restricted := make(map[string]string, len(necessary))
		for loc := range necessary {
			restricted[loc] = itemAt[loc]
		}
		g := Build(restricted, settings)
		for _, loc := range g.Order[len(g.Order)-1].locations() {
			if loc == goal {
				return true
			}
		}
		return false
	}

	for changed := true; changed; {
		changed = false
		for _, loc := range locs {
			if loc == goal || !necessary[loc] {
				continue
			}

			delete(necessary, loc)
			if reaches() {
				changed = true
			} else {
				necessary[loc] = true
			}
		}
	}

	restricted := make(map[string]string, len(necessary))
	for loc := range necessary {
		restricted[loc] = itemAt[loc]
	}

	return Build(restricted, settings)
}

package graph

import (
	"sort"
	"strings"

	"github.com/alttpr-mudora/mudora/internal/alttp/logic"
)

// VisitedSet canonically identifies a graph state: the exact set of
// progression-relevant locations checked so far, order-independent.
type VisitedSet string

func newVisitedSet(locs []string) VisitedSet {
	sorted := append([]string(nil), locs...)
	sort.Strings(sorted)
	return VisitedSet(strings.Join(sorted, "\x00"))
}

func (s VisitedSet) locations() []string {
	if s == "" {
		return nil
	}
	return strings.Split(string(s), "\x00")
}

func (s VisitedSet) with(loc string) VisitedSet {
	return newVisitedSet(append(s.locations(), loc))
}

func (s VisitedSet) withAll(locs []string) VisitedSet {
	if len(locs) == 0 {
		return s
	}
	return newVisitedSet(append(s.locations(), locs...))
}

// "Bombs (10)" is sold in several fixed, unrandomized shops; assume it's
// always reachable rather than modeling shops as locations.
var preCollectedItems = []string{"Rescue Zelda", "Bombs (10)"}

func itemsFor(s VisitedSet, itemAt map[string]string) *logic.Items {
	return itemsForLocations(s.locations(), itemAt)
}

func itemsForLocations(locs []string, itemAt map[string]string) *logic.Items {
	items := logic.NewItems()
	for _, item := range preCollectedItems {
		items.Add(item)
	}
	for _, loc := range locs {
		if item, ok := itemAt[loc]; ok {
			items.Add(item)
		}
	}
	return items
}

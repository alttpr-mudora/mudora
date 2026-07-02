package graph

import "github.com/alttpr-mudora/mudora/internal/alttp/logic"

// Edge is state -> state', labeled with every location collected in that
// step (everything reachable-and-new is always taken together).
type Edge struct {
	From, To  VisitedSet
	Locations []string
}

type Graph struct {
	ItemAt map[string]string
	Start  VisitedSet
	Order  []VisitedSet
	Edges  []Edge
}

// Build walks from the empty state, collecting every currently
// reachable-and-new location into a single combined next state each
// step. Since each step is fully determined by the current items, there's
// exactly one outgoing edge per state and the state count is bounded by
// the number of locations (each step strictly grows the visited set).
func Build(itemAt map[string]string, settings *logic.Settings) *Graph {
	itemAt = withVirtualItems(itemAt)

	var candidates []string
	for loc, item := range itemAt {
		if isProgression(item) {
			candidates = append(candidates, loc)
		}
	}

	start := newVisitedSet(nil)
	g := &Graph{ItemAt: itemAt, Start: start, Order: []VisitedSet{start}}

	state := start
	for {
		items := itemsFor(state, itemAt)
		access := regionAccessibility(items, settings)
		visited := make(map[string]bool, len(candidates))
		for _, loc := range state.locations() {
			visited[loc] = true
		}

		var batch []string
		for _, loc := range candidates {
			if !visited[loc] && locationReachable(loc, items, settings, access) {
				batch = append(batch, loc)
			}
		}

		if len(batch) == 0 {
			return g
		}

		next := state.withAll(batch)
		g.Edges = append(g.Edges, Edge{From: state, To: next, Locations: batch})
		g.Order = append(g.Order, next)
		state = next
	}
}

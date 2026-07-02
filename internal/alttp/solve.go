package alttp

import (
	"github.com/alttpr-mudora/mudora/internal/alttp/graph"
	"github.com/alttpr-mudora/mudora/internal/alttp/logic"
	"github.com/alttpr-mudora/mudora/internal/rom"
)

func Solve(data []byte) *graph.Graph {
	itemAt := make(map[string]string)
	for _, e := range rom.Inspect(data) {
		itemAt[e.Location] = e.Item
	}

	settings := logic.DefaultSettings()
	settings.LocationItems = itemAt
	medallions := rom.RequiredMedallions(data)
	settings.MiseryMireMedallion = medallions["Misery Mire"]
	settings.TurtleRockMedallion = medallions["Turtle Rock"]

	return graph.MinimalPathTo(itemAt, settings, "Ganons Tower")
}

func Reachable(data []byte) *graph.Graph {
	itemAt := make(map[string]string)
	for _, e := range rom.Inspect(data) {
		itemAt[e.Location] = e.Item
	}

	settings := logic.DefaultSettings()
	settings.LocationItems = itemAt
	medallions := rom.RequiredMedallions(data)
	settings.MiseryMireMedallion = medallions["Misery Mire"]
	settings.TurtleRockMedallion = medallions["Turtle Rock"]

	return graph.Build(itemAt, settings)
}

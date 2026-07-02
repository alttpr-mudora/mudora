package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alttpr-mudora/mudora/internal"
	"github.com/alttpr-mudora/mudora/internal/alttp/graph"
	"github.com/alttpr-mudora/mudora/internal/alttp/logic"
	"github.com/alttpr-mudora/mudora/internal/rom"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println("mudora-graph", internal.Version)
		return
	}
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: mudora-graph <rom.sfc> [out.dot]")
		fmt.Fprintln(os.Stderr, "       mudora-graph --version")
		os.Exit(2)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fail(err)
	}

	out := strings.TrimSuffix(os.Args[1], ".sfc") + ".dot"
	if len(os.Args) > 2 {
		out = os.Args[2]
	}

	itemAt := make(map[string]string)
	for _, e := range rom.Inspect(data) {
		itemAt[e.Location] = e.Item
	}

	settings := logic.DefaultSettings()
	settings.LocationItems = itemAt
	medallions := rom.RequiredMedallions(data)
	settings.MiseryMireMedallion = medallions["Misery Mire"]
	settings.TurtleRockMedallion = medallions["Turtle Rock"]

	full := graph.Build(itemAt, settings)
	shortest := graph.MinimalPathTo(itemAt, settings, "Ganons Tower")

	iconDir, err := filepath.Abs(strings.TrimSuffix(out, ".dot") + "_icons")
	if err != nil {
		fail(err)
	}
	if err := graph.WriteIcons(iconDir, full); err != nil {
		fail(err)
	}

	writeGraph(full, out, iconDir)
	shortestOut := strings.TrimSuffix(out, ".dot") + ".shortest.dot"
	writeGraph(shortest, shortestOut, iconDir)

	fmt.Fprintf(os.Stderr, "full: states=%d edges=%d\n", len(full.Order), len(full.Edges))
	fmt.Fprintf(os.Stderr, "shortest to Ganons Tower: states=%d edges=%d\n", len(shortest.Order), len(shortest.Edges))
	fmt.Fprintln(os.Stderr, "wrote", out, "and", shortestOut, "(icons in", iconDir+")")
}

func writeGraph(g *graph.Graph, out, iconDir string) {
	dotFile, err := os.Create(out)
	if err != nil {
		fail(err)
	}
	defer dotFile.Close()
	graph.WriteDOT(dotFile, g, iconDir)

	legendPath := strings.TrimSuffix(out, ".dot") + ".legend.txt"
	legendFile, err := os.Create(legendPath)
	if err != nil {
		fail(err)
	}
	defer legendFile.Close()
	graph.WriteLegend(legendFile, g)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "mudora-graph:", err)
	os.Exit(1)
}

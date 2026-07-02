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
		fmt.Fprintln(os.Stderr, "mudora-graph:", err)
		os.Exit(1)
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

	g := graph.Build(itemAt, settings)

	iconDir, err := filepath.Abs(strings.TrimSuffix(out, ".dot") + "_icons")
	if err != nil {
		fmt.Fprintln(os.Stderr, "mudora-graph:", err)
		os.Exit(1)
	}
	if err := graph.WriteIcons(iconDir, g); err != nil {
		fmt.Fprintln(os.Stderr, "mudora-graph:", err)
		os.Exit(1)
	}

	dotFile, err := os.Create(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mudora-graph:", err)
		os.Exit(1)
	}
	defer dotFile.Close()
	graph.WriteDOT(dotFile, g, iconDir)

	legendPath := strings.TrimSuffix(out, ".dot") + ".legend.txt"
	legendFile, err := os.Create(legendPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mudora-graph:", err)
		os.Exit(1)
	}
	defer legendFile.Close()
	graph.WriteLegend(legendFile, g)

	fmt.Fprintf(os.Stderr, "states=%d edges=%d\n", len(g.Order), len(g.Edges))
	fmt.Fprintln(os.Stderr, "wrote", out, legendPath, "and icons to", iconDir)
}

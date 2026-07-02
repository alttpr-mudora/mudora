package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/alttpr-mudora/mudora/internal"
	"github.com/alttpr-mudora/mudora/internal/alttp"
	"github.com/alttpr-mudora/mudora/internal/alttp/graph"
	"github.com/alttpr-mudora/mudora/internal/rom"
)

func main() {
	var romPath string
	var solve bool
	var reachable bool
	var version bool

	flag.StringVar(&romPath, "rom", "", "path to ALttPR ROM file")
	flag.StringVar(&romPath, "r", "", "path to ALttPR ROM file (shorthand)")
	flag.BoolVar(&solve, "solve", false, "print the shortest path to Ganon's Tower")
	flag.BoolVar(&solve, "s", false, "shorthand for -solve")
	flag.BoolVar(&reachable, "reachable", false, "print all reachable items")
	flag.BoolVar(&version, "version", false, "print the mudora version")
	flag.BoolVar(&version, "v", false, "shorthand for -version")
	flag.Parse()

	if version {
		fmt.Println("mudora", internal.Version)
		return
	}

	if romPath == "" {
		fmt.Fprintln(os.Stderr, "mudora: --rom is required")
		flag.Usage()
		os.Exit(2)
	}

	data, err := os.ReadFile(romPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mudora:", err)
		os.Exit(1)
	}

	if solve {
		printSolution(alttp.Solve(data))
		return
	}

	if reachable {
		printSolution(alttp.Reachable(data))
		return
	}

	query := strings.Join(flag.Args(), " ")
	for _, g := range alttp.Filter(alttp.Grouped(rom.Inspect(data)), query) {
		fmt.Println(g.Region)
		for _, p := range g.Locations {
			fmt.Printf("  %-56s %s\n", p.Location, p.Item)
		}
		fmt.Println()
	}
}

func printSolution(playthrough *graph.Graph) {
	for i, e := range playthrough.Edges {
		locs := append([]string(nil), e.Locations...)
		sort.Strings(locs)

		var plural string
		if len(locs) != 1 {
			plural = "s"
		}
		fmt.Printf("Step %d (%d check%s):\n", i+1, len(locs), plural)
		for _, loc := range locs {
			fmt.Printf("  %-56s %s\n", loc, playthrough.ItemAt[loc])
		}
	}
}

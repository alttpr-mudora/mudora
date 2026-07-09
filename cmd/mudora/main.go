package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	var readBytes bool
	var startByte string
	var byteCount int
	var romHash bool
	var permalink bool

	flag.StringVar(&romPath, "rom", "", "path to ALttPR ROM file")
	flag.StringVar(&romPath, "r", "", "path to ALttPR ROM file (shorthand)")
	flag.BoolVar(&solve, "solve", false, "print the shortest path to Ganon's Tower")
	flag.BoolVar(&solve, "s", false, "shorthand for -solve")
	flag.BoolVar(&reachable, "reachable", false, "print all reachable items")
	flag.BoolVar(&readBytes, "read-bytes", false, "read raw bytes (requires -start-byte/-byte-count)")
	flag.BoolVar(&readBytes, "rb", false, "shortand for -read-bytes")
	flag.StringVar(&startByte, "start-byte", "", "start byte for -read-bytes")
	flag.StringVar(&startByte, "sb", "", "shortand for -read-bytes")
	flag.IntVar(&byteCount, "byte-count", 0, "byte count for -read-bytes")
	flag.IntVar(&byteCount, "bc", 0, "shortand for -byte-count")
	flag.BoolVar(&romHash, "hash", false, "print ROM hash items")
	flag.BoolVar(&permalink, "permalink", false, "print the alttpr.com permalink hash embedded in the ROM")
	flag.BoolVar(&version, "version", false, "print the mudora version")
	flag.BoolVar(&version, "v", false, "shorthand for -version")
	flag.Parse()

	if version {
		fmt.Println("mudora", internal.Version)
		return
	}

	if romPath == "" {
		fmt.Fprintln(os.Stderr, "mudora: -rom is required")
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

	if readBytes {
		printBytes(data, startByte, byteCount)
		return
	}

	if romHash {
		printRomHash(data)
		return
	}

	if permalink {
		printPermalink(data)
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

func printBytes(data []byte, startByte string, byteCount int) {
	if len(startByte) == 0 {
		fmt.Fprintln(os.Stderr, "mudora: a valid hex -start-byte (-sb) value is required, ex: 0x12AB")
		os.Exit(2)
	}

	if byteCount <= 0 {
		fmt.Fprintln(os.Stderr, "mudora: -byte-count (-bc) must be > 0")
		os.Exit(2)
	}

	if strings.ToUpper(startByte[0:2]) == "0X" {
		startByte = startByte[2:]
	}

	sb, err := strconv.ParseInt(strings.ToUpper(startByte), 16, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mudora: a valid hex -start-byte (-sb) value is required, ex: 0x12AB\n\t", err)
		os.Exit(2)
	}

	bytes := rom.ReadBytes(data, sb, int64(byteCount))
	cols := 16

	for i, value := range bytes {
		if i%cols == 0 {
			fmt.Printf("0x%02x\t", sb+int64(i))
		}

		fmt.Printf("%02x", value)

		if i%cols == cols-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}

func printPermalink(data []byte) {
	hash, ok := rom.GetPermalink(data)
	if !ok {
		fmt.Fprintln(os.Stderr, "mudora: failed to retrieve permalink")
		os.Exit(1)
	}

	fmt.Println("https://alttpr.com/h/" + hash)
}

func printRomHash(data []byte) {
	items, ok := rom.GetHash(data)
	if !ok {
		fmt.Fprintln(os.Stderr, "mudora: failed to retrieve ROM hash")
		os.Exit(1)
	}

	for i, item := range items {
		fmt.Print(item)

		if i < len(items)-1 {
			fmt.Print(" | ")
		}
	}
}

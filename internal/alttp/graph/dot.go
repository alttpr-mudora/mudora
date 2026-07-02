package graph

import (
	"fmt"
	"html"
	"io"
	"path/filepath"
	"sort"

	"github.com/alttpr-mudora/mudora/internal/icons"
)

func stateIDs(g *Graph) map[VisitedSet]string {
	ids := make(map[VisitedSet]string, len(g.Order))
	for i, s := range g.Order {
		ids[s] = fmt.Sprintf("S%d", i)
	}
	return ids
}

// WriteDOT renders g as a top-to-bottom chain, keeping width bounded to
// one state box plus one location list regardless of how many states
// there are. iconDir is where those icons were written by WriteIcons and
// is used as-is in <IMG SRC>, so pass an absolute path unless you control
// the working directory `dot` gets run from.
func WriteDOT(w io.Writer, g *Graph, iconDir string) {
	ids := stateIDs(g)
	iconDir = filepath.ToSlash(iconDir)

	fmt.Fprintln(w, "digraph mudora {")
	fmt.Fprintln(w, `  bgcolor="darkgray";`)
	for _, s := range g.Order {
		fmt.Fprintf(w, "  %s [shape=box, style=filled, fillcolor=lightblue, label=%q];\n", ids[s], ids[s])
	}
	for i, e := range g.Edges {
		visit := fmt.Sprintf("V%d", i)
		fmt.Fprintf(w, "  %s [shape=plaintext, label=<%s>];\n", visit, batchTable(e.Locations, g.ItemAt, iconDir))
		fmt.Fprintf(w, "  %s -> %s;\n", ids[e.From], visit)
		fmt.Fprintf(w, "  %s -> %s;\n", visit, ids[e.To])
	}
	fmt.Fprintln(w, "}")
}

func batchTable(locs []string, itemAt map[string]string, iconDir string) string {
	var b []byte
	b = append(b, `<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="2" CELLPADDING="4" BGCOLOR="darkslategrey">`...)

	for _, loc := range locs {
		item := itemAt[loc]
		b = append(b, "<TR><TD>"...)
		if _, ok := icons.PNG(item); ok {
			b = append(b, fmt.Sprintf(`<IMG SRC="%s/%s"/>`, iconDir, iconFileName(item))...)
		}
		b = append(b, "</TD>"...)
		b = append(b, fmt.Sprintf(`<TD><FONT COLOR="white">%s</FONT></TD></TR>`, html.EscapeString(loc))...)
	}

	b = append(b, "</TABLE>"...)
	return string(b)
}

func WriteLegend(w io.Writer, g *Graph) {
	ids := stateIDs(g)
	for _, s := range g.Order {
		locs := s.locations()
		sort.Strings(locs)
		fmt.Fprintf(w, "%s (%d checks):\n", ids[s], len(locs))
		for _, loc := range locs {
			fmt.Fprintf(w, "  %s -> %s\n", loc, g.ItemAt[loc])
		}
	}
}

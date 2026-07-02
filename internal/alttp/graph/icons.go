package graph

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/alttpr-mudora/mudora/internal/icons"
)

var iconFileReplacer = strings.NewReplacer(" ", "_", "'", "", "(", "", ")", "")

func iconFileName(item string) string {
	return iconFileReplacer.Replace(item) + ".png"
}

// WriteIcons writes every item icon referenced by g's edges into dir, so
// the DOT output's <IMG SRC> references have something to point at.
func WriteIcons(dir string, g *Graph) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	written := make(map[string]bool)
	for _, e := range g.Edges {
		for _, loc := range e.Locations {
			item := g.ItemAt[loc]
			if written[item] {
				continue
			}
			written[item] = true

			data, ok := icons.PNG(item)
			if !ok {
				continue
			}
			if err := os.WriteFile(filepath.Join(dir, iconFileName(item)), data, 0o644); err != nil {
				return err
			}
		}
	}

	return nil
}

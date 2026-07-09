//go:build js && wasm

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"syscall/js"

	"github.com/alttpr-mudora/mudora/internal"
	"github.com/alttpr-mudora/mudora/internal/alttp"
	"github.com/alttpr-mudora/mudora/internal/icons"
	"github.com/alttpr-mudora/mudora/internal/rom"
)

type placement struct {
	Location    string `json:"location"`
	Item        string `json:"item"`
	Icon        string `json:"icon,omitempty"`
	Progression bool   `json:"progression"`
}

type group struct {
	Region    string      `json:"region"`
	Locations []placement `json:"locations"`
}

type itemIcon struct {
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"`
}

func main() {
	js.Global().Set("mudoraInspect", js.FuncOf(inspect))
	js.Global().Set("mudoraSolve", js.FuncOf(solve))
	js.Global().Set("mudoraItemHash", js.FuncOf(itemHash))
	js.Global().Set("mudoraPermalink", js.FuncOf(permalink))
	js.Global().Set("mudoraVersion", js.ValueOf(internal.Version))
	select {}
}

func inspect(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errResult("missing ROM bytes")
	}

	romBytes := args[0]
	data := make([]byte, romBytes.Get("length").Int())
	js.CopyBytesToGo(data, romBytes)

	query := ""
	if len(args) > 1 {
		query = args[1].String()
	}

	groups := alttp.Filter(alttp.Grouped(rom.Inspect(data)), query)
	out := make([]group, 0, len(groups))
	for _, g := range groups {
		og := group{Region: g.Region, Locations: make([]placement, 0, len(g.Locations))}
		for _, p := range g.Locations {
			pl := placement{Location: p.Location, Item: p.Item, Progression: alttp.IsProgression(p.Item)}
			if png, ok := icons.PNG(p.Item); ok {
				pl.Icon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
			}
			og.Locations = append(og.Locations, pl)
		}
		out = append(out, og)
	}

	b, err := json.Marshal(out)
	if err != nil {
		return errResult(err.Error())
	}
	return string(b)
}

func solve(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errResult("missing ROM bytes")
	}

	romBytes := args[0]
	data := make([]byte, romBytes.Get("length").Int())
	js.CopyBytesToGo(data, romBytes)

	playthrough := alttp.Solve(data)
	steps := make(map[int][]placement)

	for step, e := range playthrough.Edges {
		locs := append([]string(nil), e.Locations...)
		sort.Strings(locs)

		steps[step] = make([]placement, 0)

		for _, loc := range locs {
			item := playthrough.ItemAt[loc]
			pl := placement{Location: loc, Item: item, Progression: alttp.IsProgression(item)}
			if png, ok := icons.PNG(item); ok {
				pl.Icon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
			}
			steps[step] = append(steps[step], pl)
		}
	}

	b, err := json.Marshal(steps)
	if err != nil {
		return errResult(err.Error())
	}
	return string(b)
}

func itemHash(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errResult("missing ROM bytes")
	}

	romBytes := args[0]
	data := make([]byte, romBytes.Get("length").Int())
	js.CopyBytesToGo(data, romBytes)

	itemNames, ok := rom.GetHash(data)
	if !ok {
		return errResult("unable to resolve ROM item hash")
	}

	items := make([]itemIcon, len(itemNames))

	for i, name := range itemNames {
		ic := itemIcon{Name: name}
		if png, ok := icons.PNG(name); ok {
			ic.Icon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
		}

		items[i] = ic
	}
	b, err := json.Marshal(items)
	if err != nil {
		return errResult(err.Error())
	}
	return string(b)
}

func permalink(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errResult("missing ROM bytes")
	}

	romBytes := args[0]
	data := make([]byte, romBytes.Get("length").Int())
	js.CopyBytesToGo(data, romBytes)

	hash, ok := rom.GetPermalinkHash(data)
	if !ok {
		return errResult("unable to resolve ROM permalink")
	}

	b, err := json.Marshal(map[string]string{"permalink": fmt.Sprintf("https://alttpr.com/h/%s", hash)})
	if err != nil {
		return errResult(err.Error())
	}

	return string(b)
}

func errResult(msg string) any {
	b, _ := json.Marshal(map[string]string{"error": msg})
	return string(b)
}

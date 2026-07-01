package alttp

import "strings"

// Accessibility rules
//
// Rules are given as comma-delimited strings of items
// Items may have modifiers:
//   [item] indicates item is a logical sequence break (available, but never required to progress)
//   {item} indicates item is required to _observe_ the location, but the item is not obtainable
//   $ indicates a dynamic Lua rule
//
// [] and {} modifiers can be combined

type Rule struct {
	raw             string
	RequiredItems   []string
	IsSequenceBreak bool
}

func NewRule(raw string) *Rule {
	return &Rule{
		raw:             raw,
		RequiredItems:   formatItems(raw),
		IsSequenceBreak: strings.Index(raw, "[") > -1,
	}
}

func formatItems(rawItems string) []string {
	items := strings.Split(rawItems, ",")

	formattedItems := make([]string, 0)
	replacer := strings.NewReplacer(
		"[", "",
		"]", "",
		"{", "",
		"}", "",
	)

	for _, item := range items {
		formattedItem := replacer.Replace(item)

		if len(formattedItem) > 0 {
			formattedItems = append(formattedItems, formattedItem)
		}
	}

	return formattedItems
}

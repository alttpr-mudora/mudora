package main

import (
	"fmt"

	"github.com/alttpr-mudora/mudora/internal/alttp"
)

func main() {
	rule := alttp.NewRule("hammer,item_placement_advanced,[lamp],{}")

	fmt.Println(rule)
}

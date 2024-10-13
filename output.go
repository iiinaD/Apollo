package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
)

var BLUE_BOLD_UNDERLINE = "\033[0;1;4;34m"
var YELLOW_BOLD_UNDERLINE = "\033[0;1;4;33m"
var RESET = "\033[0m"
var GREEN = "\033[0;32m"

var renderer *glamour.TermRenderer = nil

func printResponse(text string) {
	fmt.Printf("\n%sGEMENI:%s\n", BLUE_BOLD_UNDERLINE, RESET)
	if renderer == nil {
		var err error
		renderer, err = glamour.NewTermRenderer(glamour.WithAutoStyle(), glamour.WithWordWrap(100))
		handleError(err)
	}
	out, err := renderer.Render(text)
	handleError(err)
	fmt.Println(out)
}

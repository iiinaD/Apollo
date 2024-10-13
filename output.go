package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
)

var BLUE_BOLD_UNDERLINE = "\033[0;1;4;34m"
var YELLOW_BOLD_UNDERLINE = "\033[0;1;4;33m"
var RESET = "\033[0m"
var GREEN = "\033[0;32m"

func printResponse(text string) {
	fmt.Printf("\n%sGEMENI:%s\n", BLUE_BOLD_UNDERLINE, RESET)
	r, err := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(80),
	)
	handleError(err)

	out, err := r.Render(text)
	handleError(err)
	fmt.Println(out)
}

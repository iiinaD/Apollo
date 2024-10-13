package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"github.com/google/generative-ai-go/genai"
)

var BLUE_BOLD_UNDERLINE = "\033[0;4;34m"
var RESET string = "\033[0m"

func printResponse(resp *genai.GenerateContentResponse) {
	fmt.Printf("\n%sGEMENI:%s\n", BLUE_BOLD_UNDERLINE, RESET)
	cand := resp.Candidates[0]
	text := fmt.Sprintf("%s", cand.Content.Parts[0])

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

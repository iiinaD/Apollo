package main

import (
	"context"
	"github.com/charmbracelet/glamour"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"os"
)

var client *genai.Client = nil
var model *genai.GenerativeModel = nil

func InitClient() {
	defer wg.Done()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMENI_API_KEY")))
	handleError(err)
	model = client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(getInstructions())},
	}
	theme := os.Getenv("APOLLO_OUTPUT_THEME")
	if theme != "" {
		var err error
		renderer, err = glamour.NewTermRenderer(glamour.WithStandardStyle(theme), glamour.WithWordWrap(100))
		handleError(err)
	}
}

func CloseClient() {
	if client != nil {
		client.Close()
	}
}

func getInstructions() string {
	return `You will always format the answer as markdown and always respond briefly with as short as possible explainations or answers.`
}

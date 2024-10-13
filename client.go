package main

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"os"
)

var client *genai.Client = nil
var model *genai.GenerativeModel = nil

func InitClient() {
	defer wg.Done()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	handleError(err)
	model = client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(getInstructions())},
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

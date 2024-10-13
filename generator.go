package main

import (
	"context"
	"github.com/google/generative-ai-go/genai"
)

func GenerateText(text string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	return model.GenerateContent(ctx, genai.Text(text))
}

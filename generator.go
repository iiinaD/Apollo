package main

import (
	"context"
	"github.com/google/generative-ai-go/genai"
)

func GenerateText(text string) *genai.GenerateContentResponseIterator {
	ctx := context.Background()
	return model.GenerateContentStream(ctx, genai.Text(text))
}

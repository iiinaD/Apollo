package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

var client *genai.Client = nil
var model *genai.GenerativeModel = nil

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Please provide an action for Apollo")
		return
	}
	text := os.Args[1]

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	handleError(err)
	model = client.GenerativeModel("gemini-1.5-flash")
	defer client.Close()

	res := GenerateText(text)

	readFromResponseStream(res)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readFromResponseStream(iter *genai.GenerateContentResponseIterator) {
	for {
		res, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			fmt.Println("---")
			break
		}
		handleError(err)
		printResponse(res)
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Printf("%s", part)
			}
		}
	}
}

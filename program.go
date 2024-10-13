package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var runSpinner = false
var wgSpinner = sync.WaitGroup{}

func main() {
	defer CloseClient()
	wg.Add(1)
	go InitClient()

	ctx := context.Background()

	if len(os.Args) == 2 {
		text := os.Args[1]
		wg.Wait()
		res, _ := model.GenerateContent(ctx, genai.Text(text))
		resText := getText(res)
		printResponse(resText)
		return
	}

	go spinner()

	var history []*genai.Content
	wg.Wait()
	cs := model.StartChat()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n%s\tWelcome to Apollo!%s\n\nStart chatting by writing a prompt or 'q' to exit the program:%s\n-> ", YELLOW_BOLD_UNDERLINE, GREEN, RESET)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "q" {
			break
		}
		runSpinner = true
		history = extendHistory(history, text, "user")
		cs.History = history
		resp, err := cs.SendMessage(ctx, genai.Text(text))
		handleError(err)
		respText := getText(resp)
		runSpinner = false
		wgSpinner.Wait()
		printResponse(respText)
		history = extendHistory(history, respText, "model")
		fmt.Print("-> ")
	}
}

func getText(resp *genai.GenerateContentResponse) string {
	cand := resp.Candidates[0]
	return fmt.Sprintf("%s", cand.Content.Parts[0])
}

func extendHistory(history []*genai.Content, text string, role string) []*genai.Content {
	return append(history, &genai.Content{
		Parts: []genai.Part{genai.Text(text)},
		Role:  role,
	})
}

func handleError(err error) {
	if err != nil {
		runSpinner = false
		wgSpinner.Wait()
		fmt.Println("Something happened and the program has stopped.")
		log.Fatal(err)
	}
}

func spinner() {
	for {
		if !runSpinner {
			continue
		}
		wgSpinner.Add(1)
		spinner := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			fmt.Printf("\r%s Generating response...", spinner[i%len(spinner)])
			i++
			time.Sleep(100 * time.Millisecond)
			// Replace with your condition
			if !runSpinner {
				break
			}
		}
		fmt.Printf("\r%s\r", strings.Repeat(" ", 25))
		wgSpinner.Done()
	}
}

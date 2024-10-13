package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	defer CloseClient()
	wg.Add(1)
	go InitClient()

	if len(os.Args) == 2 {
		text := os.Args[1]
		wg.Wait()
		res, _ := GenerateText(text)
		printResponse(res)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to Apollo!\n\n Start chatting by writing a prompt:\n-> ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if text == "q" {
			break
		}
		fmt.Printf("-> ")
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

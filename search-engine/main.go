package main

import (
	"bufio"
	"fmt"
	"github.com/jangirvipin/searcgh-engine/searchengine"
	"os"
	"strings"
)

func main() {
	documents := []string{
		"This is the first document.",
		"This document contains some text.",
		"This is another first with different text.",
		"This is the fourth document, which is quite different.",
		"Text is everywhere, but this document is special.",
		"First, analyze the document that contains data.",
		"Another document! Another text? Another chance.",
		"A document with just a few unique words.",
		"Different is not always better than first.",
		"This is the last document, and it contains all previous ideas.",
	}

	index := searchengine.BuildInvertedIndex(documents)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Search CLI ready! Enter your query (or 'exit'):")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		result := searchengine.Search(input, index)

		if len(result) == 0 {
			fmt.Println("No documents matched.")
			continue
		}

		fmt.Println("Matched Documents:");
		for _, idx := range result {
			fmt.Printf("Doc %d: %s\n", idx, documents[idx])
		}
	}
}

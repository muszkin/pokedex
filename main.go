package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	cleanWords := make([]string, 0)
	for _, word := range strings.Fields(text) {
		cleanWords = append(cleanWords, strings.ToLower(word))
	}
	return cleanWords
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var clicommands map[string]cliCommand

func commandHelp() error {
	fmt.Println("Usage:\n")
	for _, command := range clicommands {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	clicommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			commands := cleanInput(scanner.Text())
			command, ok := clicommands[commands[0]]
			if ok {
				if err := command.callback(); err != nil {
					fmt.Printf("something goes wrong %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	cleanWords := make([]string, 0)
	for _, word := range strings.Fields(text) {
		cleanWords = append(cleanWords, strings.ToLower(word))
	}
	return cleanWords
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

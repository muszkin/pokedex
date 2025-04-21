package main

import (
	"bufio"
	"fmt"
	poke_api "github.com/muszkin/pokedex/poke-api"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var clicommands map[string]cliCommand
var offset, limit, mapCount int
var first = true

func commandHelp() error {
	fmt.Print("Usage:\n\n")
	for _, command := range clicommands {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

func main() {
	offset = 0
	limit = 20
	mapCount = 0
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
		"map": {
			name:        "map",
			description: "Get next 20 available locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 available locations",
			callback:    commandMapb,
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

func commandMap() error {
	if !first {
		mapCount++
	} else {
		first = false
	}
	offset = mapCount * limit
	locations, err := poke_api.GetNextLocation(offset, limit)
	if err != nil {
		fmt.Printf("Something goes wrong...  %v\n", err)
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb() error {
	if mapCount > 0 {
		mapCount--
	}
	offset = mapCount * limit
	locations, err := poke_api.GetNextLocation(offset, limit)
	if err != nil {
		fmt.Printf("Something goes wrong...  %v\n", err)
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

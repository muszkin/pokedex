package main

import (
	"bufio"
	"fmt"
	poke_api "github.com/muszkin/pokedex/poke-api"
	"math/rand"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

var clicommands map[string]cliCommand
var offset, limit, mapCount int
var first = true

var pokemonStash map[string]poke_api.Pokemon

func commandHelp(...string) error {
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
	pokemonStash = map[string]poke_api.Pokemon{}
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
		"explore": {
			name:        "explore",
			description: "Explore area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "catch",
			description: "Inspect pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedec",
			description: "Show you pokemon stash",
			callback:    commandPokedex,
		},
	}
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			commands := cleanInput(scanner.Text())
			if len(commands) > 0 {
				command, ok := clicommands[commands[0]]
				if ok {
					if err := command.callback(commands[1:]...); err != nil {
						fmt.Printf("something goes wrong %v\n", err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
			continue
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

func commandExit(...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(...string) error {
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

func commandMapb(...string) error {
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

func commandExplore(args ...string) error {
	areaExploreResult, err := poke_api.ExploreLocation(args[0])
	if err != nil {
		fmt.Printf("Something goes wrong... %v\n", err)
	}
	fmt.Println("Found Pokemon:")
	for _, areaExploreResult := range areaExploreResult.PokemonEncounters {
		fmt.Println(" - " + areaExploreResult.Pokemon.Name)
	}
	return nil
}

func commandCatch(args ...string) error {
	pokemonName := args[0]
	pokemon, err := poke_api.Catch(pokemonName)
	if err != nil {
		fmt.Printf("Something goes wrong... %v\n", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catch := pokemon.BaseExperience - int(float64(pokemon.BaseExperience)*0.33)
	chance := rand.Intn(pokemon.BaseExperience)
	if chance > catch {
		fmt.Printf("%s was caught!\n\n", pokemonName)
		pokemonStash[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func commandInspect(args ...string) error {
	pokemon, ok := pokemonStash[args[0]]
	if !ok {
		fmt.Printf("You don't have this pokemon!")
		return nil
	}
	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}
	return nil
}

func commandPokedex(...string) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range pokemonStash {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}

package main

import (
	"strings"
	"fmt"
	"os"
	"math/rand"

	"github.com/CosmicRadiocity/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next string
	Previous *string
}

func cleanInput(text string) []string{
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit(cfg *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return fmt.Errorf("Unknown error exiting the program")
}

func commandHelp(cfg *config, params []string) error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, params []string) error {
	
	data, err := cfg.pokeapiClient.FetchLocationAreaBatch(cfg.Next)
	
	if err != nil { 
		return err
	}
	cfg.Next = *data.Next
	cfg.Previous = data.Previous
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config, params []string) error {
	if cfg.Previous == nil {
		fmt.Println("This is the first page.")
		return nil
	}
	data, err := cfg.pokeapiClient.FetchLocationAreaBatch(*cfg.Previous)
	if err != nil { 
		return err
	}
	cfg.Next = *data.Next
	cfg.Previous = data.Previous
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Missing parameter. Usage: explore <area name>")
	}
	name := params[0]
	fmt.Printf("Exploring %s...\n", name)
	data, err := cfg.pokeapiClient.FetchLocationAreaDetails(name)
	if err != nil{
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Missing parameter. Usage: catch <pokemon name>")
	}
	name := params[0]
	fmt.Printf("Throwing a pokeball at %s...\n", name)
	data, err := cfg.pokeapiClient.FetchPokemon(name)
	if err != nil{
		return err
	}
	

	chance := (999 - data.BaseExperience) / 15
	num :=  rand.Intn(100)
	if num <= chance {
		fmt.Printf("Caught %s!\n", name)
		cfg.pokeapiClient.AddPokemonToPokedex(name, data)
		return nil
	}

	fmt.Printf("%s escaped...\n", name)
	return nil
}

func commandInspect(cfg *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Missing parameter. Usage: inspect <pokemon name>")
	}

	pokemon, ok := cfg.pokeapiClient.GetPokemonFromPokedex(params[0])
	if !ok {
		fmt.Println("You have not caught that pokemon.")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokemonType.Type.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas",
			callback: commandMapB,
		},
		"explore": {
			name: "explore",
			description: "Displays the pokemon found in given area. Usage : explore <area name>",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch the given Pokemon. Usage : catch <pokemon name>",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Get information on a Pokemon you've caught. Usage : inspect <pokemon name>",
			callback: commandInspect,
		},
	}
}
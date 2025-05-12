package main

import (
	"strings"
	"fmt"
	"os"

	"github.com/CosmicRadiocity/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return fmt.Errorf("Unknown error exiting the program")
}

func commandHelp(cfg *config) error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("This is the first page.")
		return nil
	}
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
	}
}
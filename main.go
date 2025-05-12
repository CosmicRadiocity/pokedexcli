package main

import (
	"fmt"
	"bufio"
	"os"
	"time"

	"github.com/CosmicRadiocity/pokedexcli/internal/pokeapi"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapiClient: client,
		Next: "",
		Previous: nil,
	}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleaned := cleanInput(input)
			if cmd, ok := commands[cleaned[0]]; !ok {
				fmt.Println("Unknown command")
			} else {
				err := cmd.callback(&cfg)
				if err != nil {fmt.Printf("%v\n", err)}
			}
		}
	}
}
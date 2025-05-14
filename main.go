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
	client := pokeapi.NewClient(5 * time.Second, time.Minute*5)
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
				var params []string
				if len(cleaned) > 1 {
					params = cleaned[1:len(cleaned)]
				}
				err := cmd.callback(&cfg, params)
				if err != nil {fmt.Printf("%v\n", err)}
			}
		}
	}
}
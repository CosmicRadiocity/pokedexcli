package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
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
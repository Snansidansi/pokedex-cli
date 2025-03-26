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
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

type config struct {
	next string
	prev string
}

func getConfig() config {
	return config{
		next: "https://pokeapi.co/api/v2/location-area/?limit=20",
		prev: "",
	}
}

func startRepl() {
	commands := getCommands()
	config := getConfig()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		if cmd, ok := commands[words[0]]; ok {
			if err := cmd.callback(&config); err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Println("Unkown command")
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

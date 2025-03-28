package commands

import "github.com/snansidansi/pokedex-cli/internal/pokeapi"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "List available commands",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "List the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the specified location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch the specified pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect the specified pokemon form your pokedex",
			Callback:    CommandInspect,
		},
	}
}

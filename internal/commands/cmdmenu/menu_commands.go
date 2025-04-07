package cmdmenu

import "github.com/snansidansi/pokedex-cli/internal/commands"

func GetCommands() map[string]commands.Command {
	return map[string]commands.Command{
		"exit": {
			Name:        "exit",
			Description: "Save and exit the Pokedex",
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
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show your pokedex",
			Callback:    commandPokedex,
		},
		"reset": {
			Name:        "reset",
			Description: "Reset your progress",
			Callback:    commandReset,
		},
		"pokebox": {
			Name:        "pokebox",
			Description: "Enter the pokebox and list available operations",
			Callback:    commandPokebox,
		},
	}
}

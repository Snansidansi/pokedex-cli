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
		"pokemon": {
			Name:        "pokemon",
			Description: "List pokemon in the specified location",
			Callback:    commandPokemon,
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
			Description: "Enter the pokebox mode. Type help in this mode for more information",
			Callback:    commandPokebox,
		},
		"team": {
			Name:        "team",
			Description: "Enter the team mode. Type help in this mode for more information",
			Callback:    commandTeam,
		},
		"explore": {
			Name:        "explore",
			Description: "Start exploring",
			Callback:    commandExplore,
		},
	}
}

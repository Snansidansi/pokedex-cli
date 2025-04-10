package cmdexplore

import "github.com/snansidansi/pokedex-cli/internal/commands"

func GetCommands() map[string]commands.Command {
	return map[string]commands.Command{
		"help": {
			Name:        "help",
			Description: "List available commands",
			Callback:    commandHelp,
		},
		"back": {
			Name:        "back",
			Description: "Stop the exploration",
			Callback:    commandBack,
		},
	}
}

package cmdpokebox

import (
	"github.com/snansidansi/pokedex-cli/internal/commands"
)

func GetCommands() map[string]commands.Command {
	return map[string]commands.Command{
		"help": {
			Name:        "help",
			Description: "List available commands",
			Callback:    commandHelp,
		},
		"list": {
			Name:        "list",
			Description: "List all pokemon in the pokebox",
			Callback:    commandList,
		},
		"back": {
			Name:        "back",
			Description: "Exit the pokebox",
			Callback:    commandBack,
		},
	}
}

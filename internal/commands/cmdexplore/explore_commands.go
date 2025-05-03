package cmdexplore

import (
	"github.com/snansidansi/pokedex-cli/internal/commands"
	"github.com/snansidansi/pokedex-cli/internal/commands/basecommands"
)

func GetCommands() map[string]commands.Command {
	return map[string]commands.Command{
		"help": {
			Name:        "help",
			Description: "List available commands",
			Callback:    commandHelp,
		},
		"stop": {
			Name:        "stop",
			Description: "Stop the exploration",
			Callback:    commandStop,
		},
		"f": {
			Name:        "forward",
			Description: "Type 'f' to walk forward in the current location and you may encounter a pokemon",
			Callback:    commandForward,
		},
		"l": {
			Name:        "left",
			Description: "Type 'l' to switch to the location on your left",
			Callback:    commandLeft,
		},
		"r": {
			Name:        "right",
			Description: "Type 'r' to switch to the location on your right",
			Callback:    commandRight,
		},
		"hp": {
			Name:        "hp",
			Description: "List how much hp every pokemon in your team currently has",
			Callback:    basecommands.CommandTeamHP,
		},
		"inventory": {
			Name:        "inventory",
			Description: "List your inventory",
			Callback:    basecommands.CommandInventory,
		},
	}
}

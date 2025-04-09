package cmdteam

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
		"back": {
			Name:        "back",
			Description: "Exit the team mode",
			Callback:    commandBack,
		},
		"rename": {
			Name:        "rename",
			Description: "Rename the specified pokemon to the new given name",
			Callback:    basecommands.BaseCommandRename,
		},
		"list": {
			Name:        "list",
			Description: "List all pokemon in your team",
			Callback:    basecommands.CommandListTeam,
		},
		"list-pokebox": {
			Name:        "list-pokebox",
			Description: "List all your pokemon in the pokebox to add them to your team",
			Callback:    basecommands.CommandListPokebox,
		},
		"add": {
			Name:        "add",
			Description: "Move a pokemon from your pokebox to you team",
			Callback:    commandAdd,
		},
		"remove": {
			Name:        "remove",
			Description: "Move a pokemon from the team back to the pokebox",
			Callback:    commandRemove,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect the specified pokemon from your pokebox or team",
			Callback:    basecommands.CommandInspectOwned,
		},
	}
}

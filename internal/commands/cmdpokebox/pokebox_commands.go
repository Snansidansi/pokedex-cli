package cmdpokebox

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
		"list": {
			Name:        "list",
			Description: "List all pokemon in the pokebox",
			Callback:    basecommands.CommandListPokebox,
		},
		"back": {
			Name:        "back",
			Description: "Exit the pokebox mode",
			Callback:    commandBack,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect the specified pokemon from your pokebox or team",
			Callback:    basecommands.CommandInspectOwned,
		},
		"rename": {
			Name:        "rename",
			Description: "Rename the specified pokemon to the new given name",
			Callback:    basecommands.BaseCommandRename,
		},
		"release": {
			Name:        "release",
			Description: "Release the pokemon from your pokebox back into wild",
			Callback:    commandRelease,
		},
		"list-team": {
			Name:        "list-team",
			Description: "List all pokemon from your team",
			Callback:    basecommands.CommandListTeam,
		},
	}
}

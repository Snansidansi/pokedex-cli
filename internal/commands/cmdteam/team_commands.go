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
			Callback:    commandList,
		},
	}
}

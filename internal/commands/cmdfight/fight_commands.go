package cmdfight

import (
	"github.com/snansidansi/pokedex-cli/internal/commands"
	"github.com/snansidansi/pokedex-cli/internal/commands/basecommands"
)

func GetCommands() map[string]commands.Command {
	return map[string]commands.Command{
		"hp": {
			Name:        "hp",
			Description: "List how much hp every pokemon in your team has",
			Callback:    basecommands.CommandTeamHP,
		},
		"help": {
			Name:        "help",
			Description: "List available commands",
			Callback:    commandHelp,
		},
		"team": {
			Name:        "team",
			Description: "List all pokemon in your team",
			Callback:    basecommands.CommandListTeam,
		},
		"select": {
			Name:        "select",
			Description: "Select a pokemon from your team to fight",
			Callback:    commandSelect,
		},
	}
}

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
		"attack": {
			Name:        "attack",
			Description: "Attack the enemy pokemon with you selected pokemon",
			Callback:    commandAttack,
		},
		"stats": {
			Name:        "stats",
			Description: "Show the stats of you selected pokemon or a specified pokemon from your team",
			Callback:    commandStats,
		},
		"enemy": {
			Name:        "enemy",
			Description: "Show the stats of the enemy pokemon",
			Callback:    CommandEnemy,
		},
	}
}

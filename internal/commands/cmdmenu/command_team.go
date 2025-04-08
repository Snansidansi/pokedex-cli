package cmdmenu

import (
	"github.com/snansidansi/pokedex-cli/internal/commands/cmdteam"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandTeam(conf *pokeapi.Config, _ ...string) error {
	repl.StartRepl("Team > ", conf, cmdteam.GetCommands())
	return nil
}

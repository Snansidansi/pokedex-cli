package cmdmenu

import (
	"github.com/snansidansi/pokedex-cli/internal/commands/cmdexplore"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandExplore(conf *pokeapi.Config, _ ...string) error {
	repl.StartRepl("exploring > ", conf, cmdexplore.GetCommands())
	return nil
}

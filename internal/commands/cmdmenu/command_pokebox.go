package cmdmenu

import (
	"github.com/snansidansi/pokedex-cli/internal/commands/cmdpokebox"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandPokebox(conf *pokeapi.Config, _ ...string) error {
	repl.StartRepl("Pokebox > ", conf, cmdpokebox.GetCommands())
	return nil
}

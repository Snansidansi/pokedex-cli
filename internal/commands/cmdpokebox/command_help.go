package cmdpokebox

import (
	"github.com/snansidansi/pokedex-cli/internal/commands/basecommands"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandHelp(_ *pokeapi.Config, _ ...string) error {
	basecommands.BaseCommandHelp("Usage of the pokebox:\n\n", GetCommands())
	return nil
}

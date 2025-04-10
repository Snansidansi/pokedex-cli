package cmdmenu

import (
	"errors"

	"github.com/snansidansi/pokedex-cli/internal/commands/cmdexplore"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandExplore(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting location name or id")
	}

	locationName := args[0]
	location, err := conf.Client.GetLocation(locationName)
	if err != nil {
		return errors.New("location does not exist")
	}
	conf.CurrentLocation = location.ID

	repl.StartRepl("exploring > ", conf, cmdexplore.GetCommands())
	return nil
}

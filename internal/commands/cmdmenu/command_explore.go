package cmdmenu

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/commands/cmdexplore"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandExplore(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting location name or id")
	}

	locationNameOrID := args[0]
	location, err := conf.Client.GetLocation(locationNameOrID)
	if err != nil {
		return errors.New("location does not exist")
	}
	conf.CurrentLocationID = location.ID

	fmt.Printf("You are now in: %s\n", location.Name)
	repl.StartRepl("exploring > ", conf, cmdexplore.GetCommands())

	conf.PlayerData.Team.AfterExplorationCleanup()
	return nil
}

package cmdexplore

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandLeft(conf *pokeapi.Config, _ ...string) error {
	lowerLocationID := conf.CurrentLocationID - 1
	if lowerLocationID < 1 {
		return errors.New("Cannot go left anymore. You already are in the first location.")
	}

	conf.CurrentLocationID = lowerLocationID
	location, err := conf.Client.GetLocation(fmt.Sprint(conf.CurrentLocationID))
	if err != nil {
		return err
	}

	fmt.Printf("You are now in: %s\n", location.Name)
	return commandForward(conf)
}

package cmdexplore

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandLeft(conf *pokeapi.Config, _ ...string) error {
	conf.CurrentLocationID--
	if conf.CurrentLocationID < 1 {
		conf.CurrentLocationID++
		return errors.New("Cannot go left anymore. You already are in the first location.")
	}

	location, err := conf.Client.GetLocation(fmt.Sprint(conf.CurrentLocationID))
	if err != nil {
		conf.CurrentLocationID++
		return errors.New("The path the the next left location is to narrow.\nBut you can stop exploring and start from the location")
	}

	fmt.Printf("You are now in: %s\n", location.Name)
	return commandForward(conf)
}

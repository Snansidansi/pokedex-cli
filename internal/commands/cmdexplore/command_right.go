package cmdexplore

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandRight(conf *pokeapi.Config, _ ...string) error {
	conf.CurrentLocationID += 1

	location, err := conf.Client.GetLocation(fmt.Sprint(conf.CurrentLocationID))
	if err != nil {
		conf.CurrentLocationID -= 1
		return errors.New("Cannot go right anymore. You are already in the last location.")
	}

	fmt.Printf("You are now in: %s\n", location.Name)
	return commandForward(conf)
}

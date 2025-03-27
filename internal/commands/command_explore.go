package commands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandExplore(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting a single location: 'explore <location-name>'")
	}

	locationName := args[0]
	location, err := conf.Client.GetPokemonInLocation(locationName)
	if err != nil {
		return err
	}

	for _, encounter := range location.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

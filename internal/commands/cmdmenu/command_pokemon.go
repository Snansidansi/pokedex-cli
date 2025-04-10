package cmdmenu

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandPokemon(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting a single location: 'explore <location-name>'")
	}

	locationNameOrID := args[0]
	location, err := conf.Client.GetPokemonInLocation(locationNameOrID)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found following pokemon:")
	for _, encounter := range location.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

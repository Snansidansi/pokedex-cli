package cmdpokebox

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandInspect(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Expecting a single pokemon name from your pokebox (output from the list command)")
	}

	pokemonName := args[0]
	pokemon, ok := conf.PlayerData.Pokebox[pokemonName]
	if !ok {
		return errors.New("No pokemon with that name exists in your pokebox")
	}

	fmt.Printf("Details about %s:\n\n", pokemonName)
	pokemon.Print()

	return nil
}

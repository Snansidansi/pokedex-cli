package cmdfight

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandSelect(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting a single pokemon name")
	}

	pokemonName := args[0]
	_, ok := conf.PlayerData.Team.Get(pokemonName)
	if !ok {
		return errors.New("this pokemon is not in your team")
	}

	conf.PlayerData.Team.ActivePokemon = &pokemonName
	fmt.Printf("%s is now you active pokemon\n", pokemonName)

	return nil
}

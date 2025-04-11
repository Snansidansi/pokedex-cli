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
	pokemon, ok := conf.PlayerData.Team.Get(pokemonName)
	if !ok {
		return errors.New("this pokemon is not in your team")
	}

	if pokemon.Stats.CurrentHP <= 0 {
		return errors.New("this pokemon has no hp left")
	}

	conf.PlayerData.Team.ActivePokemon = &pokemonName
	fmt.Printf("%s is now you active pokemon\n", pokemonName)

	return nil
}

package basecommands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandHeal(conf *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("expecting at least one pokemon name")
	}

	healedPokemon := []string{}
	notExisting := []string{}
	for _, pokemonName := range args {
		err := healPokemon(conf, pokemonName)
		if err != nil {
			notExisting = append(notExisting, pokemonName)
			continue
		}
		healedPokemon = append(healedPokemon, pokemonName)
	}

	if len(healedPokemon) != 0 {
		fmt.Printf("Healed pokemon (%v):\n", len(healedPokemon))
		for _, pokemonName := range healedPokemon {
			fmt.Printf(" - %s\n", pokemonName)
		}

		if len(notExisting) != 0 {
			fmt.Println("")
		}
	}

	if len(notExisting) != 0 {
		fmt.Printf("You don't own these pokemon (%v):\n", len(notExisting))
		for _, pokemonName := range notExisting {
			fmt.Printf(" - %s\n", pokemonName)
		}
	}

	return nil
}

func healPokemon(conf *pokeapi.Config, pokemonName string) error {
	err := conf.PlayerData.Team.HealPokemon(pokemonName)
	if err == nil {
		return nil
	}

	err = conf.PlayerData.Pokebox.HealPokemon(pokemonName)
	if err == nil {
		return nil
	}

	return errors.New("pokemon does not exist int the pokebox or the team")
}

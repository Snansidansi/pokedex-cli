package commands

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandPokedex(conf *pokeapi.Config, _ ...string) error {
	if len(conf.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range conf.Pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}

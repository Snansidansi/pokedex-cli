package cmdmenu

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandPokedex(conf *pokeapi.Config, _ ...string) error {
	if conf.PlayerData.Pokedex.IsEmpty() {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemonName := range conf.PlayerData.Pokedex.GetAll() {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}

package cmdpokebox

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandHP(conf *pokeapi.Config, _ ...string) error {
	fmt.Println("Health of the pokemon in your pokebox:")

	pokemonsHP := conf.PlayerData.Pokebox.GetPokemonHPSorted()
	for _, pokemon := range pokemonsHP {
		fmt.Printf(" - %s: %v / %vhp\n", pokemon.Name, pokemon.CurrentHP, pokemon.MaxHP)
	}

	return nil
}

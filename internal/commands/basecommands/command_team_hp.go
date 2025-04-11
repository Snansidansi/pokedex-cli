package basecommands

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandTeamHP(conf *pokeapi.Config, _ ...string) error {
	fmt.Println("Health of the pokemon in your team:")

	pokemonsHP := conf.PlayerData.Team.GetPokemonHPSorted()
	for _, pokemon := range pokemonsHP {
		fmt.Printf(" - %s: %v\n", pokemon.Name, pokemon.HP)
	}

	return nil
}

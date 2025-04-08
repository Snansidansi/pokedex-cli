package basecommands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandListTeam(conf *pokeapi.Config, _ ...string) error {
	pokemonInTeam := conf.PlayerData.Team.Pokemon

	if len(pokemonInTeam) == 0 {
		return errors.New("Your team is empty")
	}

	fmt.Println("Pokemon in your team:")
	for name := range pokemonInTeam {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

package basecommands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandListTeam(conf *pokeapi.Config, _ ...string) error {
	team := conf.PlayerData.Team
	pokemonNamesSorted := team.GetAllNamesSorted()
	if len(pokemonNamesSorted) == 0 {
		return errors.New("Your team is empty")
	}

	fmt.Printf("%v of %v pokemon are in your team\n", team.Size(), team.MaxSize)
	fmt.Println("Pokemon in your team:")
	for _, name := range pokemonNamesSorted {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

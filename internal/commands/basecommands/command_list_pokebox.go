package basecommands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandListPokebox(conf *pokeapi.Config, _ ...string) error {
	pokebox := conf.PlayerData.Pokebox

	if len(pokebox) == 0 {
		return errors.New("Your pokebox is empty")
	}

	fmt.Printf("Pokemon in your pokebox (%v):\n", len(pokebox))
	for _, name := range pokebox.GetAllNamesSorted() {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

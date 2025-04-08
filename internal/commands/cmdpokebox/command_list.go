package cmdpokebox

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandList(conf *pokeapi.Config, _ ...string) error {
	pokebox := conf.PlayerData.Pokebox

	if len(pokebox) == 0 {
		return errors.New("Your pokebox is empty")
	}

	fmt.Println("Pokemon in your pokebox:")
	for _, name := range pokebox.GetAllNamesSorted() {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

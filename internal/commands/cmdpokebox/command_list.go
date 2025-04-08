package cmdpokebox

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandList(conf *pokeapi.Config, _ ...string) error {
	pokebox := conf.PlayerData.Pokebox

	if len(pokebox) == 0 {
		fmt.Println("Your pokebox is empty")
		return nil
	}

	fmt.Println("Pokemon in your pokebox:")
	for _, name := range pokebox.GetAllNamesSorted() {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

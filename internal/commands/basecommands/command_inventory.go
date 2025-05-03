package basecommands

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandInventory(conf *pokeapi.Config, _ ...string) error {
	pokeballs := entities.GetPokeballsSorted()

	fmt.Println("")
	fmt.Println("Your inventory:")
	for _, pokeball := range pokeballs {
		count := conf.PlayerData.PokeballInv[pokeball.Name]
		fmt.Printf(" - %s: %dx\n", pokeball.Name, count)
	}

	return nil
}

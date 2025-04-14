package cmdfight

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandEnemy(conf *pokeapi.Config, _ ...string) error {
	pokemon := conf.PlayerData.Team.CurrentEnemy
	asciiArt, err := conf.Client.GetAsciiImage(pokemon.ImageUrl, 15)
	if err != nil {
		return err
	}

	fmt.Print(asciiArt)
	fmt.Println("Enemy stats:")
	fmt.Printf(" - Name: %s\n", pokemon.Name)
	fmt.Printf(" - HP: %v / %vhp\n", pokemon.Stats.CurrentHP, pokemon.Stats.MaxHP)
	fmt.Printf(" - Damage: %v\n", pokemon.Stats.Damage)

	return nil
}

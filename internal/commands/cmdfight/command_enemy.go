package cmdfight

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandEnemy(conf *pokeapi.Config, _ ...string) error {
	pokemon := conf.PlayerData.Team.CurrentEnemy

	fmt.Println("Enemy stats:")
	fmt.Printf(" - Name: %s\n", pokemon.Name)
	fmt.Printf(" - HP: %v / %vhp\n", pokemon.Stats.MaxHP, pokemon.Stats.CurrentHP)
	fmt.Printf(" - Damage: %v\n", pokemon.Stats.Damage)

	return nil
}

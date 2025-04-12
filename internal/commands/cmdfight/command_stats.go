package cmdfight

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandStats(conf *pokeapi.Config, _ ...string) error {
	if conf.PlayerData.Team.ActivePokemon == nil {
		return errors.New("Your active pokemon is not in your team anymore.\nSelect a different one with the select command.")
	}

	pokemon, ok := conf.PlayerData.Team.Get(*conf.PlayerData.Team.ActivePokemon)
	if !ok {
		return errors.New("Your active pokemon is not in your team anymore.\nSelect a different one with the select command.")
	}

	fmt.Printf("Stats of %s:\n", *conf.PlayerData.Team.ActivePokemon)
	fmt.Printf(" - Pokemon: %s\n", pokemon.Name)
	fmt.Printf(" - HP: %v / %v hp\n", pokemon.Stats.CurrentHP, pokemon.Stats.MaxHP)
	fmt.Printf(" - Damage: %v\n", pokemon.Stats.Damage)
	fmt.Println("")

	return nil
}

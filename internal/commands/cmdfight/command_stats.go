package cmdfight

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandStats(conf *pokeapi.Config, args ...string) error {
	if len(args) > 1 {
		return errors.New("expecting at most one pokemon name")
	}

	var pokemonName string
	if len(args) == 0 {
		if conf.PlayerData.Team.ActivePokemon == nil {
			return errors.New("Your active pokemon is not in your team anymore.\nSelect a different one with the select command.")
		}
		pokemonName = *conf.PlayerData.Team.ActivePokemon
	} else {
		pokemonName = args[0]
	}

	pokemon, ok := conf.PlayerData.Team.Get(pokemonName)
	if !ok {
		return errors.New("this pokemon is not in your team")
	}

	asciiArt, err := conf.Client.GetAsciiImage(pokemon.ImageUrl, 15)
	if err != nil {
		return err
	}

	fmt.Print(asciiArt)
	fmt.Printf("Stats of %s:\n", pokemonName)
	fmt.Printf(" - Pokemon: %s\n", pokemon.Name)
	fmt.Printf(" - HP: %v / %v hp\n", pokemon.Stats.CurrentHP, pokemon.Stats.MaxHP)
	fmt.Printf(" - Damage: %v\n", pokemon.Stats.Damage)
	fmt.Println("")

	return nil
}

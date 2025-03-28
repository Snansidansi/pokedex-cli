package commands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandInspect(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := conf.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught %s", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Abilies:")
	for i := range pokemon.Abilities {
		fmt.Printf(" - %s\n", pokemon.Abilities[i].Ability.Name)
	}

	fmt.Println("Stats:")
	for i := range pokemon.Stats {
		stat := pokemon.Stats[i]
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for i := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon.Types[i].Type.Name)
	}

	return nil
}

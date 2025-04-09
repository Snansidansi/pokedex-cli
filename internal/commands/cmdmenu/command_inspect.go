package cmdmenu

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandInspect(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting pokemon name")
	}

	pokemonName := args[0]
	if !conf.PlayerData.Pokedex.Contains(pokemonName) {
		return fmt.Errorf("you have not caught %s", pokemonName)
	}

	pokemon, err := conf.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	printPokemonData(pokemon)

	asciiImmage, err := conf.Client.GetAsciiImage(pokemon.Sprites.Other.OfficialArtwork.FrontDefault)
	if err != nil {
		fmt.Println("image could not be displayed:")
		return err
	}
	fmt.Println("")
	fmt.Print(asciiImmage)

	return nil
}

func printPokemonData(pokemon pokeapi.PokemonDTO) {
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
}

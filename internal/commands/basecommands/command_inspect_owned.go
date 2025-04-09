package basecommands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandInspectOwned(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Expecting a single pokemon name from your pokebox (output from the list command)")
	}

	pokemonName := args[0]
	pokemon, exists := findPokemon(conf, pokemonName)

	if !exists {
		return errors.New("No pokemon with that name exists in your pokebox or team")
	}

	fmt.Printf("Details about %s:\n\n", pokemonName)
	pokemon.Print()

	asciiImage, err := conf.Client.GetAsciiImage(pokemon.ImageUrl)
	if err != nil {
		fmt.Println("image could not be displayed:")
		return err
	}
	fmt.Println("")
	fmt.Print(asciiImage)

	return nil
}

func findPokemon(conf *pokeapi.Config, pokemonName string) (entities.Pokemon, bool) {
	if pokemon, ok := conf.PlayerData.Pokebox[pokemonName]; ok {
		return pokemon, ok
	}
	if pokemon, ok := conf.PlayerData.Team.Get(pokemonName); ok {
		return pokemon, ok
	}

	return entities.Pokemon{}, false
}

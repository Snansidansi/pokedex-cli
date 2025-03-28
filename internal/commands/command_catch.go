package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandCatch(conf *pokeapi.Config, args ...string) error {
	const minCatchChance = 5

	if len(args) != 1 {
		return errors.New("expecting pokemon name")
	}

	pokemonNameOrID := args[0]
	pokemon, err := conf.Client.GetPokemon(pokemonNameOrID)
	if err != nil {
		return err
	}

	catchDifficulty := float64(pokemon.BaseExperience) / 3.85
	catchChance := 100 - int(catchDifficulty)
	catchChance = max(minCatchChance, catchChance)
	randNum := rand.Intn(101) + 1

	fmt.Printf("Chance to catch %s: %v%%\n", pokemon.Name, catchChance)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchTries := rand.Intn(4) + 3
	for range catchTries {
		fmt.Print(".")
		time.Sleep(450 * time.Millisecond)
	}
	fmt.Print(" ")

	if randNum > catchChance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	conf.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}

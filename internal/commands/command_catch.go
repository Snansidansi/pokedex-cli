package commands

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandCatch(conf *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("expecting pokemon name")
	}

	pokemonNameOrID := args[0]
	pokemon, err := conf.Client.GetPokemon(pokemonNameOrID)
	if err != nil {
		return err
	}

	pokeball := choosePokeBall(&pokemon)
	catchChance := pokemon.CalcCatchChance(pokeball)
	catched := pokemon.Catch(pokeball)

	fmt.Printf("Chance to catch %s: %v %%\n", pokemon.Name, catchChance)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchTries := rand.Intn(4) + 3
	for range catchTries {
		fmt.Print(".")
		time.Sleep(450 * time.Millisecond)
	}
	fmt.Print(" ")

	if !catched {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	conf.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}

func choosePokeBall(pokemon *pokeapi.Pokemon) pokeapi.PokeBall {
	scanner := bufio.NewScanner(os.Stdin)
	availablePokeBalls := pokeapi.GetPokeballs()

	fmt.Println("--------------------------------")
	fmt.Println("Please select a Poke Ball (type nothing for the default ball):")

	for _, pokeBall := range availablePokeBalls {
		catchChance := pokemon.CalcCatchChance(pokeBall)
		fmt.Printf(" - %s (%v %%)\n", pokeBall.Name, catchChance)
	}
	fmt.Println()

	defer fmt.Printf("--------------------------------\n\n")
	for {
		fmt.Print("Pokeball > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Print("Pokè Ball\n")
			return availablePokeBalls["Poké Ball"]
		}

		if selectedPokeBall, ok := availablePokeBalls[input]; ok {
			return selectedPokeBall
		}
		fmt.Println("This Pokeball does not exist!")
	}
}

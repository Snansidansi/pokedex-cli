package cmdmenu

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/entities/mapper"
	"github.com/snansidansi/pokedex-cli/internal/playerdata"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandCatch(conf *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("expecting pokemon name")
	}

	pokemonNameOrID := args[0]
	pokemonDTO, err := conf.Client.GetPokemon(pokemonNameOrID)
	if err != nil {
		return err
	}

	pokeball := choosePokeBall(&pokemonDTO)
	catchChance := pokemonDTO.CalcCatchChance(pokeball.CatchRateMultiplier)
	catched := pokemonDTO.Catch(pokeball)

	printCatchProcess(pokemonDTO.Name, catchChance)

	if !catched {
		fmt.Printf("%s escaped!\n", pokemonDTO.Name)
		return nil
	}

	conf.PlayerData.Pokedex.Add(pokemonDTO.Name)
	fmt.Printf("%s was caught!\n", pokemonDTO.Name)

	pokemonName := choosePokemonName(*conf.PlayerData.Pokebox, pokemonDTO.Name)
	(*conf.PlayerData.Pokebox)[pokemonName] = mapper.PokemonDTOToEntity(&pokemonDTO)

	return nil
}

func printCatchProcess(pokemonName string, catchChance int) {
	fmt.Printf("Chance to catch %s: %v %%\n", pokemonName, catchChance)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchTries := rand.Intn(4) + 3
	for range catchTries {
		fmt.Print(".")
		time.Sleep(450 * time.Millisecond)
	}
	fmt.Print(" ")
}

func choosePokemonName(pokebox playerdata.Pokebox, pokemonName string) string {
	fmt.Println("")
	fmt.Println("Name your Pokemon (type nothing for the default name):")
	fmt.Print("Name > ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputName := scanner.Text()

	fmt.Println("")

	if inputName == "" {
		return pokebox.GetNextAvailableName(pokemonName)
	}
	return pokebox.GetNextAvailableName(inputName)
}

func choosePokeBall(pokemon *pokeapi.PokemonDTO) pokeapi.PokeBall {
	scanner := bufio.NewScanner(os.Stdin)
	availablePokeBalls := pokeapi.GetPokeballs()

	fmt.Println("--------------------------------")
	fmt.Println("Please select a Poke Ball (type nothing for the default ball):")

	for _, pokeBall := range availablePokeBalls {
		catchChance := pokemon.CalcCatchChance(pokeBall.CatchRateMultiplier)
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

package basecommands

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/entities/mapper"
	"github.com/snansidansi/pokedex-cli/internal/playerdata"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandCatch(conf *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("expecting pokemon name")
	}

	pokemonNameOrID := args[0]
	pokemonDTO, err := conf.Client.GetPokemon(pokemonNameOrID)
	if err != nil {
		return err
	}

	pokeball, err := choosePokeBall(conf.PlayerData.PokeballInv, &pokemonDTO)
	if err != nil {
		return err
	}
	catchChance := pokemonDTO.CalcCatchChance(pokeball.CatchRateMultiplier)
	catched := pokemonDTO.Catch(pokeball)

	printCatchProcess(pokemonDTO.Name, catchChance)

	if !catched {
		fmt.Printf("%s escaped!\n", pokemonDTO.Name)
		return nil
	}

	conf.PlayerData.Pokedex.Add(pokemonDTO.Name)
	fmt.Printf("%s was caught!\n", pokemonDTO.Name)

	pokemonName := choosePokemonName(conf, pokemonDTO.Name)
	conf.PlayerData.Pokebox[pokemonName] = mapper.PokemonDTOToEntity(&pokemonDTO)
	fmt.Printf("%s was added to the pokebox\n", pokemonName)

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

func choosePokemonName(conf *pokeapi.Config, defaultPokemonName string) string {
	fmt.Println("")
	fmt.Println("Name your Pokemon (type nothing for the default name):")
	fmt.Print("Name > ")

	scanner := bufio.NewScanner(os.Stdin)
	pokebox := conf.PlayerData.Pokebox
	team := conf.PlayerData.Team

	scanner.Scan()
	inputName := scanner.Text()

	if inputName == "" {
		inputName = defaultPokemonName
	}

	inputName = strings.Join(strings.Fields(inputName), "-")
	assignedName := pokebox.GetNextAvailableName(inputName, team)
	return assignedName
}

func choosePokeBall(pokeballInv playerdata.PokeballInv, pokemon *pokeapi.PokemonDTO) (entities.PokeBall, error) {
	scanner := bufio.NewScanner(os.Stdin)
	allPokeballs := entities.GetPokeballsSorted()

	if pokeballInv.IsEmpty() {
		return entities.PokeBall{}, fmt.Errorf("You don't have any pokeballs.\nFleeing from %s.", pokemon.Name)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Please select a Poke Ball (type nothing for the default ball):")

	for _, pokeBall := range allPokeballs {
		amount := pokeballInv[pokeBall.Name]
		if amount < 1 {
			continue
		}

		catchChance := pokemon.CalcCatchChance(pokeBall.CatchRateMultiplier)
		fmt.Printf(" - %s (%v %%): %vx\n", pokeBall.Name, catchChance, amount)
	}
	fmt.Println()

	defer fmt.Printf("--------------------------------\n\n")
	for {
		fmt.Print("Pokeball > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("Poké Ball")
			input = "Poké Ball"
		}

		if amount, ok := pokeballInv[input]; ok {
			fmt.Printf("%s: %v\n", input, amount)
			if amount < 1 {
				fmt.Println("You do not have any pokeballs of this type left")
				continue
			}

			pokeballInv[input]--
			return entities.GetPokeballs()[input], nil
		}
		fmt.Println("This Pokeball does not exist!")
	}
}

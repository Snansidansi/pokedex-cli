package basecommands

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/playerdata"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func BaseCommandCatch(conf *pokeapi.Config, pokemon entities.Pokemon) error {
	pokeball, err := choosePokeBall(conf.PlayerData.PokeballInv, pokemon)
	if err != nil {
		return err
	}
	catchChance := pokemon.CalcCatchChance(pokeball.CatchRateMultiplier)
	catched := pokemon.Catch(pokeball)

	printCatchProcess(pokemon.Name, catchChance)

	if !catched {
		return handleNotCatched(conf, pokemon)
	}

	conf.PlayerData.Pokedex.Add(pokemon.Name)
	fmt.Printf("%s was caught!\n", pokemon.Name)

	pokemonName := choosePokemonName(conf, pokemon.Name)
	conf.PlayerData.Pokebox[pokemonName] = pokemon
	fmt.Printf("%s was added to the pokebox\n", pokemonName)

	return nil
}

func handleNotCatched(conf *pokeapi.Config, pokemon entities.Pokemon) error {
	return fmt.Errorf("%s escaped!\n", pokemon.Name)
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

func choosePokeBall(pokeballInv playerdata.PokeballInv, pokemon entities.Pokemon) (entities.PokeBall, error) {
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

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

const spacer = "--------------------------------"

func BaseCommandCatch(conf *pokeapi.Config, pokemon entities.Pokemon) error {
	for {
		pokeball, err := choosePokeBall(conf.PlayerData.PokeballInv, pokemon)
		if err != nil {
			return err
		}
		catchChance := pokemon.CalcCatchChance(pokeball.CatchRateMultiplier)
		catched := pokemon.Catch(pokeball)

		printCatchProcess(pokemon.Name, catchChance)

		if !catched {
			err := handleNotCatched(conf.PlayerData.Team, pokemon)
			if err != nil {
				return err
			}

			if !conf.PlayerData.Team.HasAliveMembers() {
				return fmt.Errorf("All pokemon in your team are defeated.\nFleeing from %s", pokemon.Name)
			}

			retry := chooseRetry()
			if retry {
				continue
			}
			return fmt.Errorf("You fled from %s", pokemon.Name)
		}

		conf.PlayerData.Pokedex.Add(pokemon.Name)
		fmt.Printf("%s was caught!\n", pokemon.Name)

		pokemonName := choosePokemonName(conf, pokemon.Name)
		conf.PlayerData.Pokebox[pokemonName] = pokemon
		fmt.Printf("%s was added to the pokebox\n", pokemonName)

		return nil
	}
}

func chooseRetry() bool {
	fmt.Println("Do you want to retry or flee?")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(retry | flee) > ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "retry":
			return true
		case "flee":
			return false
		}
	}
}

func handleNotCatched(team playerdata.Team, enemyPokemon entities.Pokemon) error {
	attackedPokemonName, died, err := team.DamageRandom(enemyPokemon.Stats.Damage)
	if err != nil {
		if team.Size() == 0 {
			return fmt.Errorf("Your team is empty and %s has no reason to stay.\n%s escaped!", enemyPokemon.Name, enemyPokemon.Name)
		}
		return fmt.Errorf("All pokemon in your team are defeated and %s has no reason to stay.\n%s escaped!", enemyPokemon.Name, enemyPokemon.Name)
	}
	teamPokemon, _ := team.Get(attackedPokemonName)

	fmt.Println("")
	fmt.Println("")
	fmt.Println(spacer)
	fmt.Printf("%s attacked your team!\n", enemyPokemon.Name)
	fmt.Printf("%s took %v damage.\n", attackedPokemonName, enemyPokemon.Stats.Damage)
	fmt.Println("")

	if !died {
		fmt.Printf("Hp of %s: %v / %vhp\n", attackedPokemonName, teamPokemon.Stats.CurrentHP, teamPokemon.Stats.MaxHP)
	} else {
		fmt.Printf("%s was defeated by %s\n", attackedPokemonName, enemyPokemon.Name)
	}
	fmt.Println("")

	randFlee := rand.Intn(100) + 1
	if randFlee > 25 {
		fmt.Printf("%s did not flee.\n", enemyPokemon.Name)
		fmt.Println(spacer)
		fmt.Println("")
		return nil
	}

	fmt.Println(spacer)
	fmt.Println("")
	return fmt.Errorf("%s escaped!\n", enemyPokemon.Name)
}

func printCatchProcess(pokemonName string, catchChance int) {
	fmt.Printf("Chance to catch %s: %v %%\n", pokemonName, catchChance)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchTries := rand.Intn(4) + 3
	for range catchTries {
		fmt.Print(".")
		time.Sleep(450 * time.Millisecond)
	}
	fmt.Println("")
	fmt.Println("")
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

	fmt.Println(spacer)
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

	defer fmt.Printf("%s\n\n", spacer)
	for {
		fmt.Print("Pokeball > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("Poké Ball")
			input = "Poké Ball"
		}

		if amount, ok := pokeballInv[input]; ok {
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

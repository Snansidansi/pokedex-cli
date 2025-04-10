package cmdexplore

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/commands/basecommands"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandForward(conf *pokeapi.Config, _ ...string) error {
	for range 3 {
		fmt.Print(".")
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Println("")

	pokemonName, ok, err := checkForEncounter(conf)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Println("Nothing here")
		return nil
	}

	fmt.Printf("You've encountered a wild %s!\n", pokemonName)
	fmt.Println("Do you want to fight or catch it?")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(catch | fight) > ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "catch":
			basecommands.CommandCatch(conf, pokemonName)
			return nil
		case "fight":
			panic("unimplmented")
			// return nil
		}
	}
}

func checkForEncounter(conf *pokeapi.Config) (string, bool, error) {
	const encounterChance = 20

	randNum := rand.Intn(101) + 1
	if randNum > encounterChance {
		return "", false, nil
	}

	location, err := conf.Client.GetLocation(fmt.Sprint(conf.CurrentLocationID))
	if err != nil {
		return "", false, err
	}

	pokemonChoice := rand.Intn(len(location.Encounters))
	pokemonName := location.Encounters[pokemonChoice].Pokemon.Name
	return pokemonName, true, nil
}

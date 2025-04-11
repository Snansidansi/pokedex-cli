package cmdexplore

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/commands/basecommands"
	"github.com/snansidansi/pokedex-cli/internal/commands/cmdfight"
	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/entities/mapper"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
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

	return manageEncounter(pokemonName, conf)
}

func manageEncounter(pokemonName string, conf *pokeapi.Config) error {
	pokemonDTO, err := conf.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemon := mapper.PokemonDTOToEntity(&pokemonDTO)

	setEnemyLevel(conf, &pokemon)

	fmt.Printf("You've encountered a wild %s level %v!\n", pokemonName, pokemon.GetLevel())
	fmt.Println("Do you want to flee or fight or catch it?")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(catch | fight | flee) > ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "catch":
			fmt.Println("")
			return basecommands.CommandCatch(conf, pokemonName)
		case "fight":
			fmt.Println("")
			return fight(conf, pokemon)
		case "flee":
			fmt.Println("")
			fmt.Printf("You fled form %s.\n", pokemonName)
			return nil
		}
	}
}

func setEnemyLevel(conf *pokeapi.Config, pokemon *entities.Pokemon) {
	avgTeamLevel := conf.PlayerData.Team.GetAverageLevel()
	pokemon.SetLevel(avgTeamLevel)
}

func fight(conf *pokeapi.Config, pokemon entities.Pokemon) error {
	conf.PlayerData.Team.CurrentEnemy = &pokemon
	defer func() { conf.PlayerData.Team.CurrentEnemy = nil }()

	repl.StartRepl("Fight > ", conf, cmdfight.GetCommands())

	if conf.PlayerData.Team.WonFight {
		fmt.Printf("You won the fight! All the Pokemon in you team gain %vxp.\n", pokemon.BaseExperience)
		conf.PlayerData.Team.AddExperience(pokemon.BaseExperience)
		return nil
	}

	fmt.Println("You lost the fight, you should heal your pokemon.")
	return commandStop(conf)
}

func checkForEncounter(conf *pokeapi.Config) (string, bool, error) {
	const encounterChance = 25

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

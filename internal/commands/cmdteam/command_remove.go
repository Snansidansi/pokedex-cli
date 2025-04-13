package cmdteam

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandRemove(conf *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("expecting pokemon names\nUsage: remove <name1> <name2> ...")
	}

	team := conf.PlayerData.Team
	pokebox := conf.PlayerData.Pokebox
	notExisting := []string{}
	removed := []string{}
	for _, name := range args {
		pokemon, ok := team.Get(name)
		if !ok {
			notExisting = append(notExisting, name)
			continue
		}

		team.Delete(name)
		pokebox[name] = pokemon
		removed = append(removed, name)
	}

	if len(removed) != 0 {
		fmt.Println("These pokemon were removed from your team:")
		for _, name := range removed {
			fmt.Printf(" - %s\n", name)
		}

		if len(notExisting) != 0 {
			fmt.Println("")
		}
	}

	if len(notExisting) != 0 {
		fmt.Println("Could not remove pokemon from the team because they don't exist in your team:")
		for _, name := range notExisting {
			fmt.Printf(" - %s\n", name)
		}
	}

	return nil
}

package cmdteam

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandAdd(conf *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("expecting pokemon names\nUsage: add <name1> <name2> ...")
	}

	team := conf.PlayerData.Team
	emptyTeamSlots := int(team.MaxSize) - team.Size()

	if len(args) > emptyTeamSlots {
		return fmt.Errorf("you do not have enough empty slots in your team.\nEmpty slots: %v\nMax slots: %v\n", emptyTeamSlots, team.MaxSize)
	}

	pokebox := conf.PlayerData.Pokebox
	notExisting := []string{}
	added := []string{}
	for _, name := range args {
		pokemon, ok := pokebox[name]
		if !ok {
			notExisting = append(notExisting, name)
			continue
		}

		delete(pokebox, name)
		team.Add(name, pokemon)
		added = append(added, name)
	}

	if len(added) != 0 {
		fmt.Println("These pokemon were added to your team:")
		for _, name := range added {
			fmt.Printf(" - %s\n", name)
		}

		if len(notExisting) != 0 {
			fmt.Println("")
		}
	}

	if len(notExisting) != 0 {
		fmt.Println("Could not add pokemon to the pokebox because they dont exist in your pokebox:")
		for _, name := range notExisting {
			fmt.Printf(" - %s\n", name)
		}
	}

	return nil
}

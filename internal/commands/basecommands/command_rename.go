package basecommands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandRename(conf *pokeapi.Config, args ...string) error {
	if len(args) != 2 {
		return errors.New("Wrong usage of Command.\nUsage: rename <old-name> <new-name>")
	}

	oldName, newName := args[0], args[1]
	if oldName == newName {
		return nil
	}

	pokebox := conf.PlayerData.Pokebox
	if _, ok := pokebox[newName]; ok {
		return errors.New("pokemon with the new name already exists in the pokebox")
	}

	team := conf.PlayerData.Team
	if _, ok := team.Get(newName); ok {
		return errors.New("pokemon with the new name already exists in the team")
	}

	if pokemon, ok := pokebox[oldName]; ok {
		delete(pokebox, oldName)
		pokebox[newName] = pokemon
	} else if _, ok := team.Get(oldName); ok {
		team.Rename(oldName, newName)
	} else {
		return errors.New("you do not have a pokemon with the specified name")
	}

	fmt.Printf("%s was successfuly renamed to %s\n", oldName, newName)
	return nil
}

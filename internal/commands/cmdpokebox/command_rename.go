package cmdpokebox

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandRename(conf *pokeapi.Config, args ...string) error {
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
	if _, ok := team.Pokemon[newName]; ok {
		return errors.New("pokemon with the new name already exists in the team")
	}

	pokemon := pokebox[oldName]
	delete(pokebox, oldName)
	pokebox[newName] = pokemon

	fmt.Printf("%s was successfuly renamed to %s\n", oldName, newName)
	return nil
}

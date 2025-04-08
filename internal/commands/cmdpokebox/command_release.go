package cmdpokebox

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandRelease(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expecting a single pokemon name")
	}

	pokemonName := args[0]
	pokebox := *conf.PlayerData.Pokebox
	if _, ok := pokebox[pokemonName]; !ok {
		return errors.New("pokemon with this name does not exist in your pokebox")
	}

	fmt.Printf("Are you sure to release %s?\n", pokemonName)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(yes | no) > ")
		scanner.Scan()
		confirmDelete := scanner.Text()

		switch confirmDelete {
		case "yes":
			delete(pokebox, pokemonName)
			fmt.Printf("%s was successfuly released into the wild\n", pokemonName)
			return nil
		case "no":
			fmt.Printf("Aborting release of %s\n", pokemonName)
			return nil
		}
	}
}

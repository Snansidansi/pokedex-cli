package commands

import (
	"fmt"
	"os"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func CommandExit(_ *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

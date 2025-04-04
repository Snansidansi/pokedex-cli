package cmdmenu

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandHelp(_ *pokeapi.Config, _ ...string) error {
	fmt.Println("")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println("---------------------------------------------------")
	return nil
}

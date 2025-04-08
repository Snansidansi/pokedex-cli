package basecommands

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/commands"
)

func BaseCommandHelp(headerMessage string, commands map[string]commands.Command) {
	fmt.Println("")
	fmt.Println("---------------------------------------------------")
	fmt.Print(headerMessage)

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println("---------------------------------------------------")
}

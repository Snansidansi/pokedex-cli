package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/snansidansi/pokedex-cli/internal/commands"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func StartRepl(promptMessage string, config *pokeapi.Config, commands map[string]commands.Command) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(promptMessage)
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var commandArgs []string = nil
		if len(words) > 1 {
			commandArgs = words[1:]
		}

		if cmd, ok := commands[commandName]; ok {
			err := cmd.Callback(config, commandArgs...)
			if errors.Is(err, ExitReplError{}) {
				return
			}

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("")
			continue
		}

		fmt.Println("Unkown command")
		fmt.Println("")
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

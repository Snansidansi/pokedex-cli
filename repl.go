package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/snansidansi/pokedex-cli/internal/commands/cmdmenu"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func startRepl(config *pokeapi.Config) {
	commands := cmdmenu.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
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
			if err := cmd.Callback(config, commandArgs...); err != nil {
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
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

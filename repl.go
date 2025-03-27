package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/snansidansi/pokedex-cli/internal/commands"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func startRepl(config *pokeapi.Config) {
	commands := commands.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		if cmd, ok := commands[words[0]]; ok {
			if err := cmd.Callback(config); err != nil {
				fmt.Println(err)
			}
			fmt.Println()
			continue
		}

		fmt.Println("Unkown command")
		fmt.Println()
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

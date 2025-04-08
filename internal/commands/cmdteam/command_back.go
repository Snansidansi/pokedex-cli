package cmdteam

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandBack(_ *pokeapi.Config, _ ...string) error {
	fmt.Println("Leaving the team mode")
	return repl.ExitReplError{}
}

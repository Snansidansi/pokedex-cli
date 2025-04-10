package cmdexplore

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

func commandBack(_ *pokeapi.Config, _ ...string) error {
	fmt.Println("Stopping the exploration")
	return repl.ExitReplError{}
}

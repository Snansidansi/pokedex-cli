package commands

import "github.com/snansidansi/pokedex-cli/internal/pokeapi"

type Command struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, ...string) error
}

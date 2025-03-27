package main

import (
	"time"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func main() {
	config := pokeapi.Config{
		Client: pokeapi.NewClient(5 * time.Second),
	}

	startRepl(&config)
}

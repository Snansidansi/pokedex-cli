package main

import (
	"fmt"
	"os"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/playerdata"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	config := pokeapi.Config{
		Client:     pokeapi.NewClient(5*time.Second, 2*time.Minute, 5*time.Minute),
		PlayerData: playerdata.NewPlayerData(),
		SaveFile: pokeapi.SaveFile{
			Dir: workingDir,
		},
	}

	if err := config.Load(); err != nil {
		fmt.Println("Error while reading savefile.")
		fmt.Println(err)
		os.Exit(1)
	}

	startRepl(&config)
}

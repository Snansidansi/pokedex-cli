package cmdmenu

import (
	"fmt"
	"os"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandExit(conf *pokeapi.Config, _ ...string) error {
	const maxRetrys = 3
	fmt.Println("Start saving data")

	var err error = nil
	for range maxRetrys {
		if err = conf.Save(); err != nil {
			fmt.Println(err)
			fmt.Println("")
			fmt.Println("Retrying...")
		}
	}
	if err == nil {
		fmt.Println("Data saved successfuly")
	}

	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

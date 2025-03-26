package commands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandMapb(conf *pokeapi.Config) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMapb")
	}

	if conf.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := pokeapi.Getlocations(conf.Prev)
	if err != nil {
		return err
	}

	conf.Next = *locationAreas.Next
	if locationAreas.Previous == nil {
		conf.Prev = ""
	} else {
		conf.Prev = *locationAreas.Previous
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	fmt.Println("")

	return nil
}

package commands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandMap(conf *pokeapi.Config) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMap")
	}

	if conf.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	locationAreas, err := pokeapi.Getlocations(conf.Next)
	if err != nil {
		return err
	}

	if locationAreas.Next == nil {
		conf.Next = ""
	} else {
		conf.Next = *locationAreas.Next
	}

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

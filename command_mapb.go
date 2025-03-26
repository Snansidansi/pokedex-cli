package main

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/pokeapi"
)

func commandMapb(conf *config) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMapb")
	}

	if conf.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := pokeapi.Getlocations(conf.prev)
	if err != nil {
		return err
	}

	conf.next = *locationAreas.Next
	if locationAreas.Previous == nil {
		conf.prev = ""
	} else {
		conf.prev = *locationAreas.Previous
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	fmt.Println("")

	return nil
}

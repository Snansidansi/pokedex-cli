package main

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/pokeapi"
)

func commandMap(conf *config) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMap")
	}

	if conf.next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	locationAreas, err := pokeapi.Getlocations(conf.next)
	if err != nil {
		return err
	}

	if locationAreas.Next == nil {
		conf.next = ""
	} else {
		conf.next = *locationAreas.Next
	}

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

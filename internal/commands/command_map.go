package commands

import (
	"errors"
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandMap(conf *pokeapi.Config, _ ...string) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMap")
	}

	location, err := conf.Client.GetLocations(conf.NextLocationURL)
	if err != nil {
		return err
	}

	conf.NextLocationURL = location.Next
	conf.PrevLocationURL = location.Previous

	for _, location := range location.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(conf *pokeapi.Config, _ ...string) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMapb")
	}

	if conf.PrevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := conf.Client.GetLocations(conf.PrevLocationURL)
	if err != nil {
		return err
	}

	conf.NextLocationURL = locations.Next
	conf.PrevLocationURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

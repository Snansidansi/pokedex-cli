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

	locationAreas, err := conf.Client.GetLocationAreas(conf.NextLocationURL)
	if err != nil {
		return err
	}

	conf.NextLocationURL = locationAreas.Next
	conf.PrevLocationURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(conf *pokeapi.Config) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMapb")
	}

	if conf.PrevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreas, err := conf.Client.GetLocationAreas(conf.PrevLocationURL)
	if err != nil {
		return err
	}

	conf.NextLocationURL = locationAreas.Next
	conf.PrevLocationURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

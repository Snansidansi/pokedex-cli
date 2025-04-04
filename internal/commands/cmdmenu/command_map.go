package cmdmenu

import (
	"errors"
	"fmt"
	"strings"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandMap(conf *pokeapi.Config, _ ...string) error {
	if conf == nil {
		return errors.New("The config pointer is nil - commandMap")
	}

	locations, err := conf.Client.GetLocations(conf.NextLocationURL)
	if err != nil {
		return err
	}

	conf.NextLocationURL = locations.Next
	conf.PrevLocationURL = locations.Previous

	printLocations(locations)

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

	printLocations(locations)

	return nil
}

func printLocations(locations pokeapi.Locations) {
	for _, location := range locations.Results {
		splitURL := strings.Split(location.URL, "/")
		id := splitURL[len(splitURL)-2]
		fmt.Printf("%v - %s\n", id, location.Name)
	}
}

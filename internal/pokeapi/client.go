package pokeapi

import (
	"encoding/json"
	"net/http"
)

type Config struct {
	Next string
	Prev string
}

func GetConfig() Config {
	return Config{
		Next: "https://pokeapi.co/api/v2/location-area/?limit=20",
		Prev: "",
	}
}

func Getlocations(url string) (location_areas, error) {
	res, err := http.Get(url)
	if err != nil {
		return location_areas{}, err
	}
	defer res.Body.Close()

	locationAreas := location_areas{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return location_areas{}, err
	}

	return locationAreas, nil
}

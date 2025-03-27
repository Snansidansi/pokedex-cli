package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInLocation(location_name string) (Location, error) {
	url := baseURL + "/location-area/" + location_name

	data, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Location{}, err
		}
	}

	pokemonInLocation := Location{}
	if err := json.Unmarshal(data, &pokemonInLocation); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return pokemonInLocation, nil
}

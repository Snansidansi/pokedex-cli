package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInLocation(locationNameOrID string) (Location, error) {
	url := baseURL + "/location-area/" + locationNameOrID

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

		if resp.StatusCode != http.StatusOK {
			return Location{}, errors.New("location does not exist")
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Location{}, err
		}
	}

	pokemonInLocation := Location{}
	if err := json.Unmarshal(data, &pokemonInLocation); err != nil {
		return Location{}, err
	}

	cacheData, err := json.Marshal(pokemonInLocation)
	if err == nil {
		c.cache.Add(url, cacheData)
	}

	return pokemonInLocation, nil
}

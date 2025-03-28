package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonNameOrID string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonNameOrID

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return Pokemon{}, errors.New("pokemon does not exist")
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	cacheData, err := json.Marshal(pokemon)
	if err == nil {
		c.cache.Add(url, cacheData)
	}

	return pokemon, nil
}

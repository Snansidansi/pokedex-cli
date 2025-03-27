package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(inputURL *string) (Locations, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if inputURL != nil {
		url = *inputURL
	}

	data, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Locations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Locations{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Locations{}, err
		}
	}

	locations := Locations{}
	if err := json.Unmarshal(data, &locations); err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, data)
	return locations, nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(inputURL *string) (Location_areas, error) {
	url := baseURL + "/location-area?limit=20"
	if inputURL != nil {
		url = *inputURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location_areas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location_areas{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location_areas{}, err
	}

	locationAreas := Location_areas{}
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return Location_areas{}, err
	}

	return locationAreas, nil
}

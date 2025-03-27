package pokeapi

import (
	"encoding/json"
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

	locationAreas := Location_areas{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return Location_areas{}, err
	}

	return locationAreas, nil
}

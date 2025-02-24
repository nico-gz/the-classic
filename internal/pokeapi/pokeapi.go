package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (client *Client) GetLocations(pageUrl *string) (Location, error) {
	url := baseURL + "/location-area"
	var locations Location
	if pageUrl != nil {
		url = *pageUrl
	}

	// Go to cache before requesting
	if data, ok := client.cache.Get(url); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return locations, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}
	res, err := client.httpClient.Do(req)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}

	if res.StatusCode > 299 {
		return locations, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, data)
	}

	if err = json.Unmarshal(data, &locations); err != nil {
		return locations, err
	}

	client.cache.Add(url, data)
	return locations, nil
}

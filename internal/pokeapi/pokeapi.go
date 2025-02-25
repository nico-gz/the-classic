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

type AreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func (client *Client) GetPokemonInArea(location string) ([]string, error) {
	if location == "" {
		return nil, fmt.Errorf("error: no location name provided")
	}
	url := baseURL + "/location-area/" + location
	var areaData AreaData
	var encounters []string

	// Go to cache before requesting
	if data, ok := client.cache.Get(url); ok {
		if err := json.Unmarshal(data, &areaData); err != nil {
			return encounters, err
		}
		for _, encounter := range areaData.PokemonEncounters {
			encounters = append(encounters, encounter.Pokemon.Name)
		}
		return encounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return encounters, err
	}
	res, err := client.httpClient.Do(req)
	if err != nil {
		return encounters, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return encounters, err
	}

	if res.StatusCode > 299 {
		return encounters, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, data)
	}

	if err = json.Unmarshal(data, &areaData); err != nil {
		return encounters, err
	}

	for _, encounter := range areaData.PokemonEncounters {
		encounters = append(encounters, encounter.Pokemon.Name)
	}

	client.cache.Add(url, data)
	return encounters, nil
}

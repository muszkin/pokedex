package poke_api

import (
	"encoding/json"
	"fmt"
	poke_cache "github.com/muszkin/pokedex/poke-cache"
	"io"
	"net/http"
	"time"
)

type LocationAreaResults struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type LocationAreaExploreResult struct {
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

var duration, _ = time.ParseDuration("60s")
var cache = poke_cache.NewCache(duration)

const locationAreaUrl = "https://pokeapi.co/api/v2/location-area/"

func GetNextLocation(offset, limit int) (LocationAreaResults, error) {
	var locations LocationAreaResults
	url := fmt.Sprintf("%s?offset=%d&limit=%d", locationAreaUrl, offset, limit)
	data, inCache := cache.Get(url)
	if inCache {
		if err := json.Unmarshal(data, &locations); err != nil {
			return locations, err
		}
		return locations, nil
	}
	request, err := http.Get(url)
	if err != nil {
		return locations, err
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return locations, err
	}
	if err := json.Unmarshal(body, &locations); err != nil {
		return locations, err
	}
	defer cache.Add(url, body)
	return locations, nil
}

func ExploreLocation(area string) (LocationAreaExploreResult, error) {
	var exploreAreaResults LocationAreaExploreResult
	url := fmt.Sprintf("%s%s", locationAreaUrl, area)
	data, inCache := cache.Get(url)
	if inCache {
		if err := json.Unmarshal(data, &exploreAreaResults); err != nil {
			return exploreAreaResults, err
		}
		return exploreAreaResults, nil
	}
	request, err := http.Get(url)
	if err != nil {
		return exploreAreaResults, err
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return exploreAreaResults, err
	}
	if err := json.Unmarshal(body, &exploreAreaResults); err != nil {
		return exploreAreaResults, err
	}
	defer cache.Add(url, body)
	return exploreAreaResults, nil
}

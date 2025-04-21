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

package main

import (
	"fmt"
	"io"
	"net/http"
	pokecache "pokedexcli/internal/pokecache"
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func fetchData(endpoint string, cache *pokecache.Cache) ([]byte, error) {
	if data, exists := cache.Get(endpoint); exists {
			return data, nil
	}
	res, err := http.Get(endpoint)
	if err != nil {
			return nil, fmt.Errorf("error fetching data from API: %v", err)
	}
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
	}
	
	cache.Add(endpoint, body)
	
	return body, nil
}
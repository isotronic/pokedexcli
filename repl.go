package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	pokecache "pokedexcli/internal/pokecache"
	"strings"
)

type CLICommand struct {
	name string
	description string
	callback func(config *ConfigType) error
}

type ConfigType struct {
	nextEndpoint string
	previousEndpoint string
	cache *pokecache.Cache
}

type mapResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]CLICommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *ConfigType) error {
	endpoint := "https://pokeapi.co/api/v2/location-area/"
	if config.nextEndpoint != "" {
		endpoint = config.nextEndpoint
	}

	body, err := fetchData(endpoint, config.cache)
	if err != nil {
		return err
	}

	var data mapResult
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	config.nextEndpoint = data.Next
	if data.Previous != nil {
		config.previousEndpoint = *data.Previous
	} else {
		config.previousEndpoint = ""
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(config *ConfigType) error {
	endpoint := ""
	if config.previousEndpoint != "" {
		endpoint = config.previousEndpoint
	} else {
		fmt.Println("You're on the first page")
		return nil
	}

	body, err := fetchData(endpoint, config.cache)
	if err != nil {
		return err
	}

	var data mapResult
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	config.nextEndpoint = data.Next
	if data.Previous != nil {
		config.previousEndpoint = *data.Previous
	} else {
		config.previousEndpoint = ""
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
)

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

	var data MapResult
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

	var data MapResult
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

func commandExplore(config *ConfigType) error {
	endpoint := "https://pokeapi.co/api/v2/location-area/"

	if len(config.arg) == 0 {
		return fmt.Errorf("you need to add a location name")
	}

	endpoint += config.arg
	body, err := fetchData(endpoint, config.cache)
	if err != nil {
		return err
	}

	var data ExploreResult
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	fmt.Printf("Exploring %v...\n", config.arg)
	fmt.Println("Found Pokemon:")
	for _, encounters := range data.PokemonEncounters {
		fmt.Printf(" - %v\n", encounters.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *ConfigType) error {
	endpoint := "https://pokeapi.co/api/v2/pokemon/"

	if len(config.arg) == 0 {
		return fmt.Errorf("you need to add a pokemon name")
	}

	endpoint += config.arg
	body, err := fetchData(endpoint, config.cache)
	if err != nil {
		return err
	}

	var result PokemonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

		fmt.Printf("Throwing a Pokeball at %v...\n", result.Name)

	catchProbability := 5 / math.Pow(math.Log(float64(result.BaseExperience)), 1.4)
	if rand.Float64() <= catchProbability {
		fmt.Printf("%v was caught!\n", result.Name)
		config.pokedex[result.Name] = result
	} else {
		fmt.Printf("%v escaped!\n", result.Name)
	}

	return nil
}

func commandInspect(config *ConfigType) error {
	if len(config.arg) == 0 {
		return fmt.Errorf("you need to add a pokemon name")
	}

	pokemon, exists := config.pokedex[config.arg]
	if !exists {
		return fmt.Errorf("you have not caught %v yet", config.arg)
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %v\n", pokemonType.Type.Name)
	}

	return nil
}
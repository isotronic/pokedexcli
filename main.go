package main

import (
	"bufio"
	"fmt"
	"os"
	pokecache "pokedexcli/internal/pokecache"
	"time"
)

func main() {
	commands := make(map[string]CLICommand)
	config := ConfigType{}
	commands["help"] = CLICommand{
		name: "help",
		description: "Displays a help message",
		callback: func(config *ConfigType) error {
			return commandHelp(commands)
		},
	}
	commands["map"] = CLICommand{
		name: "map",
		description: "List 20 Pokemon location areas",
		callback: func(config *ConfigType) error {
			return commandMap(config)
		},
	}
	commands["mapb"] = CLICommand{
		name: "mapb",
		description: "List the previous 20 Pokemon location areas",
		callback: func(config *ConfigType) error {
			return commandMapb(config)
		},
	}
	commands["exit"] = CLICommand{
			name: "exit",
			description: "Exit the Pokedex",
			callback: func(config *ConfigType) error {
				return commandExit()
			},
	}

	config.cache = pokecache.NewCache(10 * time.Second)


	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		cleaned := words[0]

		command, ok := commands[cleaned]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		
		err := command.callback(&config)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
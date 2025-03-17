package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := make(map[string]cliCommand)
	config := configType{}
	commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: func(config *configType) error {
			return commandHelp(commands)
		},
	}
	commands["map"] = cliCommand{
		name: "map",
		description: "List 20 Pokemon location areas",
		callback: func(config *configType) error {
			return commandMap(config)
		},
	}
	commands["exit"] = cliCommand{
			name: "exit",
			description: "Exit the Pokedex",
			callback: func(config *configType) error {
				return commandExit()
			},
	}


	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)[0]

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
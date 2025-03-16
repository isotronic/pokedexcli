package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := make(map[string]cliCommand)
	commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: func() error {
			return commandHelp(commands)
		},
	}
	commands["exit"] = cliCommand{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
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
		
		err := command.callback()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
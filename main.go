package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//REPL loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		//converts input into slice of lowercase words
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		//isolate first word as intended command
		command := words[0]

		//Use command to find valid commands from a map and run its callback if it is found
		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		//if command is not valid
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:		 "help",
			description: "Display a help message",
			callback:	 commandHelp,
		},
		"exit": {
			name:		 "exit",
			description: "Exit the Pokedex",
			callback:	 commandExit,
		},
	}
}
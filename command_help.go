package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex\nUsage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
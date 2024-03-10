package main

import "fmt"

func callBackHelp(cfg *config) error {
	availableCommands := cliCommands()
	fmt.Println("Available commands:")
	for _, command := range availableCommands {
		fmt.Println("\t", command.name, "-", command.description)
	}
	return nil
}

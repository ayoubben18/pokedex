package main

import "fmt"

func callBackHelp() {
	availableCommands := cliCommands()
	fmt.Println("Available commands:")
	for _, command := range availableCommands {
		fmt.Println("\t", command.name, "-", command.description)
	}
}

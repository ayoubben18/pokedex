package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	// code at the bottom to a loop
	for {
		fmt.Print("> ")
		//scan and get the text
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		availableCommands := cliCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("unknown command: ", commandName)
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

// func returns map of strings to cli command
func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    callBackExit,
		},
		"map": {
			name:        "map",
			description: "Lists some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "map back",
			description: "Lists previous location areas",
			callback:    callbackMapB,
		},
		"help": {
			name:        "help",
			description: "show available commands",
			callback:    callBackHelp,
		},
	}
}

// clean input func takes a string and return a slice of strings
func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}

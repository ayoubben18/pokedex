package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

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
			continue
		}
		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

// func returns map of strings to cli command
func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    callBackExit,
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*CommandConfig, ...string) error
}

func runRepl(config *CommandConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		userInput := cleanInput(input)
		commandName := userInput[0]
		commandArgs := []string{}
		if len(userInput) > 1 {
			commandArgs = userInput[1:]
		}

		if command, ok := getCommands()[commandName]; ok {
			err := command.callback(config, commandArgs...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Println("Unknown command")

	}

	//fmt.Printf("read line: %s-\n", scanner.Text())
}

func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	cleanText := strings.Fields(lowercased)
	return cleanText
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View data about a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check the Pokemon you captured",
			callback:    commandPokedex,
		},
	}
}

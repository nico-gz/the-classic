package main

import (
	"fmt"
	"os"

	"github.com/nico-gz/pokedexcli/internal/pokeapi"
)

type CommandConfig struct {
	PokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func commandExit(config *CommandConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *CommandConfig) error {
	locations, err := config.PokeapiClient.GetLocations(config.Next)
	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(config *CommandConfig) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := config.PokeapiClient.GetLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandHelp(config *CommandConfig) error {
	fmt.Println("Welcome to the Pokedex!\nUsage")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

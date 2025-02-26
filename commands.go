package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/nico-gz/pokedexcli/internal/pokeapi"
)

type CommandConfig struct {
	PokeapiClient pokeapi.Client
	Pokedex       map[string]pokeapi.Pokemon
	Next          *string
	Previous      *string
	Id            *string
}

func commandExit(config *CommandConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *CommandConfig, args ...string) error {
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

func commandMapb(config *CommandConfig, args ...string) error {
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

func commandHelp(config *CommandConfig, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

// TODO: review args indexing being hardcoded constants
func commandExplore(config *CommandConfig, args ...string) error {
	area := args[0]
	if area == "" {
		return fmt.Errorf("missing required argument <area-name>")
	}
	encounters, err := config.PokeapiClient.GetPokemonInArea(area)
	if err != nil {
		return err
	}
	for _, encounter := range encounters {
		fmt.Println(encounter)
	}
	return nil
}

func commandCatch(config *CommandConfig, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		return fmt.Errorf("missing required argument <pokemon-name>")
	}
	pokemon, err := config.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if 36+rand.Intn(580) < pokemon.BaseExperience {
		fmt.Printf("%s escaped\n", pokemonName)
	} else {
		fmt.Printf("%s was captured\n", pokemonName)
		config.Pokedex[pokemonName] = pokemon
	}

	return nil
}

func commandInspect(config *CommandConfig, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		return fmt.Errorf("missing required argument <pokemon-name>")
	}
	pokemon, ok := config.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("that pokemon has not been caught yet")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weigth: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

	return nil
}

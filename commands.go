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

/*
After a user uses the map commands to find a location area, we want them to be able to see a list of all the Pokémon located there.
 Tips

    Use the same PokeAPI location-area endpoint, but this time you'll need to pass the name of the location area being explored. By adding a name or id, the API will return a lot more information about the location area.
    Feel free to use tools like JSON lint and JSON to Go to help you parse the response.
    Parse the Pokemon's names from the response and display them to the user.
    Make sure to use the caching layer again! Re-exploring an area should be blazingly fast.
    You'll need to alter the function signature of all your commands to allow them to allow parameters. E.g. explore <area_name>

*/

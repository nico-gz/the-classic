package main

import (
	"time"

	"github.com/nico-gz/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 30*time.Second)
	config := &CommandConfig{
		PokeapiClient: pokeClient,
	}
	runRepl(config)
}

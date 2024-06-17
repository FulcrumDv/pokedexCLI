package main

import "github.com/brexxel/pokedexCLI/internal/pokeapi"

func main() {
	pokeapiClient := pokeapi.NewClient()

	cfg := &config{
		pokeapiClient: pokeapiClient,
	}

	startRepl(cfg)
}

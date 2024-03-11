package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location area provided")
	}
	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:", locationArea.Name)
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Println(" - ", pokemonEncounter.Pokemon.Name)
	}

	return nil
}

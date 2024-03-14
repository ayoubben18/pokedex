package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("Pokemon %s was not caught yet", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("- %s\n", ability.Ability.Name)
	}

	return nil
}

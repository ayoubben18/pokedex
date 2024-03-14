package main

import "fmt"

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokedex:")
	fmt.Println("--------")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Base experience: %d\n", pokemon.BaseExperience)
		fmt.Println("--------")

	}

	return nil
}

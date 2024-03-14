package main

import (
	"github.com/ayoubben18/pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

// we are building a REPL read evaluate print loop
func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)

}

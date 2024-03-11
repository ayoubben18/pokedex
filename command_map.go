package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, locationArea := range resp.Results {
		fmt.Println(" - ", locationArea.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous
	return nil
}

func callbackMapB(cfg *config, args ...string) error {
	if cfg.previousLocationAreaUrl == nil {
		return errors.New("No previous location area")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, locationArea := range resp.Results {
		fmt.Println(" - ", locationArea.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous
	return nil
}

package main

import (
	"fmt"

	"github.com/Itssanthoshhere/pokedex-cli/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

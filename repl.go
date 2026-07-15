package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

		// switch command {
		// case "help":
		// 	fmt.Println("Welcome to Pokedex help menu!")
		// 	fmt.Println("Here are your available commands:")
		// 	fmt.Println(" - help")
		// 	fmt.Println(" - exit")
		// 	fmt.Println("")
		// case "exit":
		// 	os.Exit(0)

		// default:
		// 	fmt.Println("Invalid command")
		// }

		// fmt.Println("echoing: ", cleaned)
	}
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "Prints the help menu",
			callback: callbackHelp,
		},
		"map": {
			name:     "map",
			desc:     "Lists the next page of location areas",
			callback: callbackMap,
		},
		"mapb": {
			name:     "maps",
			desc:     "Lists the previous page of  location areas",
			callback: callbackMapb,
		},
		"explore": {
			name:     "explore {location_area}",
			desc:     "Lists the pokemon in a location area",
			callback: callbackExplore,
		},
		"catch": {
			name:     "catch {pokemon_name}",
			desc:     "Attempt to catch a pokemon and add it to your pokedex",
			callback: callbackCatch,
		},
		"exit": {
			name:     "exit",
			desc:     "Turns off the Pokedex",
			callback: callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

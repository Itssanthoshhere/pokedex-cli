package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Println("Error:", err)
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
	callback func() error
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
			desc:     "Lists some location areas",
			callback: callbackMap,
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

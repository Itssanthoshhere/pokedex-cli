# Pokedex CLI Setup Guide

Welcome to the Pokedex CLI project! This is a command-line interface application built in Go that interacts with the [PokeAPI](https://pokeapi.co/).

## Installation

1. Make sure you have [Go](https://go.dev/) installed.
2. Navigate to your project directory.
3. Run the following command to build the program:

```bash
go build
```

## Running the CLI

After building the program, you can run the executable:

```bash
./pokedex-cli
```

You will see the Pokedex REPL prompt ` >` where you can type your commands.

## Available Commands

Here are the commands you can use in the Pokedex CLI:

* **`help`**: Displays the help menu with all available commands.
* **`exit`**: Exits the Pokedex REPL.
* **`map`**: Displays the next 20 location areas in the Pokémon world. Successive calls will paginate through the list.
* **`mapb`**: Displays the previous 20 location areas.
* **`explore {location_area}`**: Lists all the Pokémon that can be found in the specified location area. (e.g., `explore mt-coronet-2f`)
* **`catch {pokemon_name}`**: Attempts to catch a Pokémon by its name. The success rate depends on the Pokémon's base experience. (e.g., `catch machoke`)
* **`inspect {pokemon_name}`**: Shows the stats (Height, Weight, Stats, and Types) of a Pokémon that you have already caught. If you haven't caught it yet, you won't be able to inspect it!

## Running Tests

To run the internal tests (such as the cache behavior), you can use:

```bash
go test ./...
```

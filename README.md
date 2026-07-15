# Pokedex CLI 🐉

A fully-featured Command Line Interface (CLI) application built in Go that lets you explore the Pokémon world right from your terminal! It interacts with the [PokeAPI](https://pokeapi.co/) to fetch locations, explore areas, catch Pokémon, and manage your very own Pokedex.

## Features ✨

- **Interactive REPL**: A built-in command loop that takes your inputs and provides instantaneous feedback.
- **Smart Caching**: Implements an in-memory caching system to minimize API calls to PokeAPI, making repeated commands lightning fast.
- **Exploration & Catching**: Discover different location areas in the Pokémon world, see what Pokémon live there, and try to catch them! Catch rates are calculated based on the Pokémon's base experience.
- **Personal Pokedex**: Keep track of the Pokémon you've caught and inspect their detailed stats (Height, Weight, Stats, and Types).

## Installation 🛠️

1. Ensure you have [Go](https://go.dev/) installed on your machine.
2. Clone this repository (or navigate to your project directory):
   ```bash
   cd pokedex-cli
   ```
3. Build the application:
   ```bash
   go build
   ```

## Usage 🚀

Run the compiled executable to start the Pokedex REPL:

```bash
./pokedex-cli
```

You will be greeted with the ` >` prompt. Type any of the available commands to interact with the pokedex!

### Commands

- `help` - Prints the help menu with all available commands.
- `exit` - Safely exits the Pokedex REPL.
- `map` - Lists the next 20 location areas in the Pokémon world. Successive calls will paginate forward through the list.
- `mapb` - Lists the previous 20 location areas (paginates backward).
- `explore {location_area}` - Lists all the Pokémon that can be found in a specific location area.
  - _Example:_ `explore mt-coronet-2f`
- `catch {pokemon_name}` - Throws a Pokeball at a Pokémon! The catch success rate depends on the Pokémon's base experience.
  - _Example:_ `catch machoke`
- `inspect {pokemon_name}` - Shows the details and base stats of a Pokémon. **Note:** You can only inspect Pokémon you have successfully caught!
  - _Example:_ `inspect machoke`
- `pokedex` - Displays a list of all the Pokémon you currently have in your Pokedex.

## Example Session 🎮

```text
 >map
Location areas:
 - canalave-city-area
 - eterna-city-area
...
 >explore mt-coronet-2f
Exploring mt-coronet-2f...
Found Pokemon:
 - clefairy
 - golbat
 - machoke
...
 >catch machoke
machoke was caught!
 >pokedex
Pokemon in Pokedex
 - machoke
 >inspect machoke
Name machoke
Height 15
Weight 705
Stats:
  - hp: 80
  - attack: 100
  - defense: 70
...
Types:
  - fighting
 >exit
```

## Running Tests 🧪

To verify the internal behavior (like the caching logic), run the Go tests:

```bash
go test ./...
```

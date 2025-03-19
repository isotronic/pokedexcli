# Pokedex CLI

Pokedex CLI is a command-line interface application that allows users to explore, catch, and inspect Pokémon using data from the PokéAPI.

## Features

- **help**: Displays a help message with available commands.
- **map**: Lists 20 Pokémon location areas.
- **mapb**: Lists the previous 20 Pokémon location areas.
- **explore**: Shows all the Pokémon that can be encountered in a given location area.
- **catch**: Tries to catch a Pokémon and add it to your Pokedex.
- **inspect**: Shows details for a Pokémon you already caught.
- **pokedex**: Shows all the Pokémon you already caught.
- **exit**: Exits the Pokedex CLI.

## Installation

1. Clone the repository:

   ```sh

   git clone https://github.com/isotronic/pokedexcli.git
   cd pokedexcli
   ```

## Usage

Run the application:

```sh
go run .
```

You will see a prompt where you can enter commands. For example:

```sh
Pokedex > help
```

## Commands

- help: Displays a help message with available commands.
- map: Lists 20 Pokémon location areas.
- mapb: Lists the previous 20 Pokémon location areas.
- explore [location]: Shows all the Pokémon that can be encountered in a given location area.
- catch [pokemon]: Tries to catch a Pokémon and add it to your Pokedex.
- inspect [pokemon]: Shows details for a Pokémon you already caught.
- pokedex: Shows all the Pokémon you already caught.
- exit: Exits the Pokedex CLI.

## Testing

Run the tests:

```sh
go test ./...
```

## License

This project is licensed under the MIT License.

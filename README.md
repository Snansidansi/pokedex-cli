# pokedex-cli
Basic cli pokedex repl written in go that uses the online pokeapi.

## How to use
1. Clone the github repo and run the the program with `go run .` (in the project root) or build the project.
2. Type help to see the available commands

## Features
- list the different locations from the pokeapi
- explore a location (see what pokemon are in that location)
- try to catch a pokemon with different pokeballs
- caught pokemon apper in the pokedex and in the pokebox
- add pokemon from your pokebox to your team and level them over time or in fights
- playerdata (pokedex, pokebox, ...) are saved to a file when exiting with the `exit` command
- All requests to the pokeapi get cached some time to reduce traffic

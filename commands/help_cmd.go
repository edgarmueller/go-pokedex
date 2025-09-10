package commands

import (
	"fmt"

	"github.com/edgarmueller/go-pokedex/internal"
)

func Help(opts []string, g *internal.Game) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

map: Displays a map of the current location
mapb: Displays a map of the current location
explore <location>: Explore the current location
catch <pokemon>: Catch a Pokemon
inspect <pokemon>: Inspect a Pokemon
pokedex: Display the pokedex
help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

package commands

import (
	"fmt"

	"github.com/edgarmueller/go-pokedex/internal"
)

func Pokedex(opts []string, g *internal.Game) error {
	fmt.Println("Your Pokedex")
	for _, p := range g.Pokedex {
		fmt.Println(" - " + p.Name)
	}
	return nil
}

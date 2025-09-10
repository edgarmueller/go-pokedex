package commands

import (
	"errors"
	"fmt"

	"github.com/edgarmueller/go-pokedex/internal"
)

func Catch(opts []string, g *internal.Game) error {
	if len(opts) != 1 {
		return errors.New("no pokemon provided")
	}
	p, err := internal.RequestPokemon(opts[0])
	if err != nil {
		return errors.New("Pokemon not found: " + err.Error())
	}
	catched := g.AttemptCatch(p)
	if catched {
		fmt.Printf("You catched the %s!\n", p.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("The %s escaped!\n", p.Name)
	}
	return nil
}

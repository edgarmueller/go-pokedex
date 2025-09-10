package commands

import (
	"errors"
	"fmt"

	"github.com/edgarmueller/go-pokedex/internal"
)

func Inspect(opts []string, g *internal.Game) error {
	if len(opts) != 1 {
		return errors.New("no pokemon provided")
	}
	idOrName := opts[0]
	p, err := g.GetPokemon(idOrName)

	if err != nil {
		return err
	}

	fmt.Println("Name: " + p.Name)
	fmt.Println("Height: " + fmt.Sprint(p.Height))
	fmt.Println("Weight: " + fmt.Sprint(p.Weight))
	fmt.Println("Stats: ")
	for _, s := range p.Stats {
		fmt.Println("- " + s.Stat.Name + ": " + fmt.Sprint(s.BaseStat))
	}
	fmt.Println("Types: ")
	for _, t := range p.Types {
		fmt.Println("- " + t.Type.Name)
	}
	return nil
}

package commands

import (
	"fmt"

	"github.com/edgarmueller/go-pokedex/internal"
)

func MapForwards(opts []string, g *internal.Game) error {
	locations, err := g.MoveToNextLocationAreas()
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}

	return nil
}

func MapBackwards(opts []string, g *internal.Game) error {
	locations, err := g.MoveToPrevLocationAreas()
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}

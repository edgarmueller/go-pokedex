package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/edgarmueller/go-pokedex/internal"
)

func Explore(opts []string, g *internal.Game) error {
	if len(opts) != 1 {
		return errors.New("no location provided")
	}
	idOrName := opts[0]
	log.Println("Exploring location " + idOrName)
	areaData, err := g.GetLocationArea(idOrName)
	if err != nil {
		return err
	}
	log.Println("Found Pokemon:")
	for _, p := range areaData.PokemonEncounters {
		fmt.Println("- " + p.Pokemon.Name)
	}
	return nil
}

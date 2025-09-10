package commands

import (
	"fmt"
	"os"

	"github.com/edgarmueller/go-pokedex/internal"
)

func Exit(opts []string, g *internal.Game) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

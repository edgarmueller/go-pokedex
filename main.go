package main

import (
	"github.com/edgarmueller/go-pokedex/internal"
)

func main() {
	g := internal.NewGame()
	startRepl(g)
}

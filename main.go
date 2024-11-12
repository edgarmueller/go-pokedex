package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/edgarmueller/go-pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(opts []string, gs *internal.Game) error
}

func commandHelp(opts []string, g *internal.Game) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:
map: Displays a map of the current location
mapb: Displays a map of the current location
explore <location>: Explore the current location
help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandExit(opts []string, g *internal.Game) error {
	os.Exit(0)
	return nil
}

func commandMap(opts []string, g *internal.Game) error {
	locations, err := g.GetNextLocationAreas()
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapB(opts []string, g *internal.Game) error {
	locations, err := g.GetPrevLocationAreas()
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}

func commandExplore(opts []string, g *internal.Game) error {
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

func commandCatch(opts []string, g *internal.Game) error {
	p, err := internal.RequestPokemon(opts[0])
	if err != nil {
		return errors.New("Pokemon not found: " + err.Error())
	}
	catched := g.AttemptCatch(p)
	if catched {
		fmt.Println("You catched the Pokemon!")
	} else {
		fmt.Println("The Pokemon escaped!")
	}
	return nil
}

func commandInspect(opts []string, g *internal.Game) error {
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

func commandPokedex(opts []string, g *internal.Game) error {
	fmt.Println("Your Pokedex")
	for _, p := range g.Pokedex {
		fmt.Println(" - " + p.Name)
	}
	return nil
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Displays a map of the current location",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays a map of the current location",
		callback:    commandMapB,
	},
	"explore": {
		name:        "explore",
		description: "Explore the current location",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a Pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a Pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Display the pokedex",
		callback:    commandPokedex,
	},
}

func main() {
	g := internal.NewGame()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("pokedex> ")
		scanner.Scan()
		cmdLine := scanner.Text()
		cmdArray := strings.Split(cmdLine, " ")
		c, exists := commands[cmdArray[0]]
		if !exists {
			fmt.Println("Command not found")
		} else {
			err := c.callback(cmdArray[1:], g)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/edgarmueller/go-pokedex/commands"
	"github.com/edgarmueller/go-pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(opts []string, gs *internal.Game) error
}

func cleanInput(text string) []string {
	re := regexp.MustCompile(`\S+`)
	tokens := re.FindAllString(strings.TrimSpace(text), -1)
	return tokens
}

func startRepl(g *internal.Game) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error or EOF: %v\n", err)
			break
		}
		cmdArray := cleanInput(line)
		cmdName := cmdArray[0]
		c, exists := getCommands()[cmdName]
		if !exists {
			fmt.Println("Command not found: " + cmdName)
			os.Exit(1)
		} else {
			err := c.callback(cmdArray[1:], g)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commands.Help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commands.Exit,
		},
		"map": {
			name:        "map",
			description: "Displays a map of the current location",
			callback:    commands.MapForwards,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a map of the current location",
			callback:    commands.MapBackwards,
		},
		"explore": {
			name:        "explore",
			description: "Explore the current location",
			callback:    commands.Explore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commands.Catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			callback:    commands.Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display the pokedex",
			callback:    commands.Pokedex,
		},
	}
}

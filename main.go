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

func cleanInput(text string) []string {
	re := regexp.MustCompile(`\S+`)
	tokens := re.FindAllString(strings.TrimSpace(text), -1)
	return tokens
}

type cliCommand struct {
	name        string
	description string
	callback    func(opts []string, gs *internal.Game) error
}

var cmds = map[string]cliCommand{
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

func main() {
	g := internal.NewGame()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error or EOF: %v\n", err)
			break
		}
		cmdArray := cleanInput(line)
		c, exists := cmds[cmdArray[0]]
		if !exists {
			fmt.Println("Command not found: " + cmdArray[0])
			os.Exit(1)
		} else {
			err := c.callback(cmdArray[1:], g)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}
	}
}

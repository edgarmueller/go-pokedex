package internal

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/edgarmueller/go-pokedex/internal/pokecache"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Game struct {
	currentPage int
	cache       *pokecache.Cache
	Pokedex     map[string]Pokemon
}

type Pokemon = structs.Pokemon

func NewGame() *Game {
	return &Game{
		currentPage: 0,
		cache:       pokecache.NewCache(5 * time.Second),
		Pokedex:     make(map[string]Pokemon),
	}
}

func (gs *Game) MoveToNextLocationAreas() ([]structs.Result, error) {
	gs.currentPage += 1
	locations, err := RequestLocationAreas(gs.currentPage, gs.cache)

	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (gs *Game) MoveToPrevLocationAreas() ([]structs.Result, error) {
	gs.currentPage -= 1
	if gs.currentPage < 0 {
		gs.currentPage = 1
	}
	locations, err := RequestLocationAreas(gs.currentPage, gs.cache)

	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (gs *Game) GetLocationArea(idOrName string) (structs.LocationArea, error) {
	key := "location-area-" + idOrName
	l, ok := gs.cache.Get(key)

	if ok {
		var area structs.LocationArea
		err := json.Unmarshal(l, &area)
		if err != nil {
			return structs.LocationArea{}, err
		}
		return area, nil
	} else {
		areas, err := pokeapi.LocationArea(idOrName)
		if err != nil {
			return structs.LocationArea{}, err
		}
		data, err := json.Marshal(areas)
		if err != nil {
			return structs.LocationArea{}, err
		}
		gs.cache.Add(key, data)
		return areas, nil
	}
}

func (g *Game) AttemptCatch(p Pokemon) bool {
	// Calculate catch probability
	catchProbability := calculateCatchProbability(p.BaseExperience)

	// Generate a random float between 0 and 1
	attempt := rand.Float64()

	// Catch the Pokemon if the random attempt is less than the catch probability
	catched := attempt < catchProbability
	if catched {
		// Add the Pokemon to the list of catched Pokemon
		g.Pokedex[p.Name] = p
		return true
	}

	return false
}

func (g *Game) GetPokemon(idOrName string) (Pokemon, error) {
	p, ok := g.Pokedex[idOrName]
	if !ok {
		return Pokemon{}, errors.New("Pokemon not found")
	}
	return p, nil
}

func calculateCatchProbability(baseExp int) float64 {
	// reference value highest base experience, see https://pwo-wiki.info/index.php/Base_Experience
	// no idea whether this is accurate
	maxBaseExp := 210
	minCatchChance := 0.1 // min chance of catching
	maxCatchChance := 0.7 // max chance of catching

	// Linear scale the baseExp into the range of catch chance
	catchChance := maxCatchChance - (float64(baseExp)/float64(maxBaseExp))*(maxCatchChance-minCatchChance)

	// Ensure catchChance stays within [minCatchChance, maxCatchChance]
	if catchChance < minCatchChance {
		catchChance = minCatchChance
	} else if catchChance > maxCatchChance {
		catchChance = maxCatchChance
	}

	return catchChance
}

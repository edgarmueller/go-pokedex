package internal

import (
	"encoding/json"
	"strconv"

	"github.com/edgarmueller/go-pokedex/internal/pokecache"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

const limit = 20

func RequestLocationAreas(page int, cache *pokecache.Cache) ([]structs.Result, error) {
	p := strconv.Itoa(page)
	key := "location-area-page-" + p
	l, ok := cache.Get(key)

	if ok {
		var locations []structs.Result
		err := json.Unmarshal(l, &locations)
		if err != nil {
			return []structs.Result{}, err
		}
		return locations, nil
	} else {
		l, err := pokeapi.Resource("location-area", page*limit, limit)
		if err != nil {
			return []structs.Result{}, err
		}
		data, err := json.Marshal(l.Results)
		if err != nil {
			return nil, err
		}
		cache.Add(key, data)
		return l.Results, nil
	}
}

func RequestLocationArea(idOrName string, cache *pokecache.Cache) (structs.LocationArea, error) {
	key := "location-area-" + idOrName
	l, ok := cache.Get(key)

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
		cache.Add(key, data)
		return areas, nil
	}
}

func RequestPokemon(idOrName string) (Pokemon, error) {
	p, err := pokeapi.Pokemon(idOrName)
	if err != nil {
		return Pokemon{}, err
	}
	return p, nil
}

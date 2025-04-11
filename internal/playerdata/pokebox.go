package playerdata

import (
	"errors"
	"fmt"
	"slices"

	"github.com/snansidansi/pokedex-cli/internal/entities"
)

type Pokebox map[string]entities.Pokemon

func (p Pokebox) GetNextAvailableName(pokemonName string, team Team) string {
	_, existsInPokebox := p[pokemonName]
	_, existsInTeam := team.Get(pokemonName)
	if !existsInPokebox && !existsInTeam {
		return pokemonName
	}

	i := 2
	for {
		name := fmt.Sprintf("%s%v", pokemonName, i)
		if _, ok := p[name]; ok {
			i++
			continue
		}
		return name
	}
}

func (p Pokebox) GetAllNamesSorted() []string {
	sortedPokemon := make([]string, len(p))
	i := 0
	for givenName := range p {
		sortedPokemon[i] = givenName
		i++
	}

	slices.Sort(sortedPokemon)

	return sortedPokemon
}

func (p Pokebox) GetPokemonHPSorted() []entities.PokemonHP {
	pokemonsHP := make([]entities.PokemonHP, len(p))
	i := 0
	for name, pokemon := range p {
		pokemonsHP[i] = entities.PokemonHP{
			Name:      name,
			CurrentHP: pokemon.Stats.CurrentHP,
			MaxHP:     pokemon.Stats.MaxHP,
		}
		i++
	}

	slices.SortFunc(pokemonsHP, entities.SortPokemonHP)
	return pokemonsHP
}

func (p Pokebox) HealPokemon(pokemonName string) error {
	pokemon, ok := p[pokemonName]
	if !ok {
		return errors.New("pokemon does not exists in the pokebox")
	}

	pokemon.Stats.CurrentHP = pokemon.Stats.MaxHP
	p[pokemonName] = pokemon

	return nil
}

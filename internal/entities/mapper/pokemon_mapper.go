package mapper

import (
	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func PokemonDTOToEntity(pokemonDTO *pokeapi.PokemonDTO) entities.Pokemon {
	stats, baseStats := mapStats(pokemonDTO)
	types := mapTypes(pokemonDTO)

	return entities.Pokemon{
		ID:                pokemonDTO.ID,
		Name:              pokemonDTO.Name,
		BaseExperience:    pokemonDTO.BaseExperience,
		CurrentExperience: 0,
		Height:            pokemonDTO.Height,
		Weight:            pokemonDTO.Weight,
		Stats:             stats,
		BaseStats:         baseStats,
		Types:             types,
		ImageUrl:          pokemonDTO.Sprites.Other.OfficialArtwork.FrontDefault,
	}
}

func mapStats(pokemonDTO *pokeapi.PokemonDTO) (stats, baseStats entities.Stats) {
	tempStats := make(map[string]int, len(pokemonDTO.Stats))
	for _, stat := range pokemonDTO.Stats {
		tempStats[stat.Stat.Name] = stat.BaseStat
	}

	stats = entities.Stats{
		MaxHP:     tempStats["hp"],
		CurrentHP: tempStats["hp"],
		Damage:    tempStats["attack"],
	}

	baseStats = entities.Stats{
		MaxHP:     stats.MaxHP,
		CurrentHP: 0,
		Damage:    stats.Damage,
	}

	return stats, baseStats
}

func mapTypes(pokemonDTO *pokeapi.PokemonDTO) []string {
	types := make([]string, len(pokemonDTO.Types))
	i := 0
	for _, t := range pokemonDTO.Types {
		types[i] = t.Type.Name
		i++
	}

	return types
}

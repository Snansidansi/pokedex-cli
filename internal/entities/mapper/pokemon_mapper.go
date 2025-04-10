package mapper

import (
	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func PokemonDTOToEntity(pokemonDTO *pokeapi.PokemonDTO) entities.Pokemon {
	stats, baseStats := mapStats(pokemonDTO)

	return entities.Pokemon{
		ID:                pokemonDTO.ID,
		Name:              pokemonDTO.Name,
		BaseExperience:    pokemonDTO.BaseExperience,
		CurrentExperience: 0,
		Height:            pokemonDTO.Height,
		Order:             pokemonDTO.Order,
		Weight:            pokemonDTO.Weight,
		Abilities:         pokemonDTO.Abilities,
		Stats:             stats,
		BaseStats:         baseStats,
		Forms:             pokemonDTO.Forms,
		Types:             pokemonDTO.Types,
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
		Speed:     tempStats["speed"],
	}

	baseStats = entities.Stats{
		MaxHP:     stats.MaxHP,
		CurrentHP: 0,
		Damage:    stats.Damage,
		Speed:     stats.Speed,
	}

	return stats, baseStats
}

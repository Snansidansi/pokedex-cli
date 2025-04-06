package mapper

import (
	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func PokemonDTOToEntity(pokemonDTO *pokeapi.PokemonDTO) entities.Pokemon {
	return entities.Pokemon{
		ID:                pokemonDTO.ID,
		Name:              pokemonDTO.Name,
		BaseExperience:    pokemonDTO.BaseExperience,
		CurrentExperience: 0,
		Height:            pokemonDTO.Height,
		Order:             pokemonDTO.Order,
		Weight:            pokemonDTO.Weight,
		Abilities:         pokemonDTO.Abilities,
		Forms:             pokemonDTO.Forms,
		Types:             pokemonDTO.Types,
	}
}

package playerdata

import (
	"fmt"

	"github.com/snansidansi/pokedex-cli/internal/entities"
)

type Team struct {
	Pokemon map[string]entities.Pokemon
	MaxSize uint
}

func NewTeam(maxSize uint) Team {
	return Team{
		Pokemon: map[string]entities.Pokemon{},
		MaxSize: maxSize,
	}
}

func (team Team) Add(pokemonName string, pokemon entities.Pokemon) error {
	if len(team.Pokemon) >= int(team.MaxSize) {
		return fmt.Errorf("team is full (max size: %v)\n", team.MaxSize)
	}

	if _, ok := team.Pokemon[pokemonName]; ok {
		return fmt.Errorf("pokemon with that name already exists in team: %s", pokemonName)
	}

	team.Pokemon[pokemonName] = pokemon

	return nil
}

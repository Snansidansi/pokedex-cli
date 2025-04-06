package playerdata_test

import (
	"testing"

	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/playerdata"
)

func TestPokebox_GetDefaultName(t *testing.T) {
	defaultPokebox := playerdata.Pokebox{
		"pikachu1":     entities.Pokemon{},
		"pikachu2":     entities.Pokemon{},
		"namedPokemon": entities.Pokemon{},
		"bulbasaur1":   entities.Pokemon{},
		"charmander1":  entities.Pokemon{},
		"charmander3":  entities.Pokemon{},
	}
	cases := []struct {
		name        string
		pokebox     playerdata.Pokebox
		pokemonName string
		expected    string
	}{
		{
			name:        "Get next name with empty pokebox",
			pokebox:     playerdata.Pokebox{},
			pokemonName: "charmander",
			expected:    "charmander1",
		},
		{
			name:        "Get next name with not contained name and filled pokebox",
			pokebox:     defaultPokebox,
			pokemonName: "squirtle",
			expected:    "squirtle1",
		},
		{
			name:        "Get next name with pokemon already existing once",
			pokebox:     defaultPokebox,
			pokemonName: "bulbasaur",
			expected:    "bulbasaur2",
		},
		{
			name:        "Get next name with pokemon already existing twice",
			pokebox:     defaultPokebox,
			pokemonName: "pikachu",
			expected:    "pikachu3",
		},
		{
			name:        "name1 and name3 already exist. Should return name2",
			pokebox:     defaultPokebox,
			pokemonName: "charmander",
			expected:    "charmander2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.pokebox.GetDefaultName(c.pokemonName)
			if actual != c.expected {
				t.Errorf("GetDefaultName() = %v, want %v", actual, c.expected)
				return
			}
		})
	}
}

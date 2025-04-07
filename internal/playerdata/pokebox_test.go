package playerdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/snansidansi/pokedex-cli/internal/entities"
)

func TestPokebox_GetDefaultName(t *testing.T) {
	defaultPokebox := Pokebox{
		"pikachu":      entities.Pokemon{},
		"pikachu2":     entities.Pokemon{},
		"namedPokemon": entities.Pokemon{},
		"bulbasaur":    entities.Pokemon{},
		"charmander":   entities.Pokemon{},
		"charmander3":  entities.Pokemon{},
	}
	cases := []struct {
		name        string
		pokebox     Pokebox
		pokemonName string
		expected    string
	}{
		{
			name:        "Get next name with empty pokebox",
			pokebox:     Pokebox{},
			pokemonName: "charmander",
			expected:    "charmander",
		},
		{
			name:        "Get next name with not contained name and filled pokebox",
			pokebox:     defaultPokebox,
			pokemonName: "squirtle",
			expected:    "squirtle",
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
			actual := c.pokebox.GetNextAvailableName(c.pokemonName)
			if actual != c.expected {
				t.Errorf("GetDefaultName() = %v, want %v", actual, c.expected)
				return
			}
		})
	}
}

func TestGetAllNamesSorted(t *testing.T) {
	cases := []struct {
		name     string
		pokebox  Pokebox
		expected []string
	}{
		{
			name:     "empty pokebox",
			pokebox:  Pokebox{},
			expected: []string{},
		},
		{
			name: "not empty pokebox",
			pokebox: Pokebox{
				"a": entities.Pokemon{},
				"b": entities.Pokemon{},
				"c": entities.Pokemon{},
			},
			expected: []string{"a", "b", "c"},
		},
		{
			name: "not empty pokebox with numbers",
			pokebox: Pokebox{
				"a": entities.Pokemon{},
				"z": entities.Pokemon{},
				"1": entities.Pokemon{},
				"9": entities.Pokemon{},
			},
			expected: []string{"1", "9", "a", "z"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.pokebox.GetAllNamesSorted()
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetAllNamesSorted() mismatch (-want + got):\n%s", diff)
				return
			}
		})
	}
}

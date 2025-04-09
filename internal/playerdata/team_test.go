package playerdata

import (
	"testing"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/entities"
)

func TestNewTeam(t *testing.T) {
	cases := []struct {
		name     string
		maxsize  uint
		expected Team
	}{
		{
			name:    "positive maxsize",
			maxsize: 5,
			expected: Team{
				Pokemon: map[string]entities.Pokemon{},
				MaxSize: 5,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := NewTeam(c.maxsize, 0, 0)

			if len(actual.Pokemon) != 0 {
				t.Errorf("NewTeam() did not create an empty map of pokemon: %v\n", len(actual.Pokemon))
			}

			if actual.MaxSize != c.expected.MaxSize {
				t.Errorf("NewTeam() does not have expected maxsize.\nExpected: %v\nActual: %v\n", c.expected.MaxSize, actual.MaxSize)
			}
		})
	}
}

func TestTeamAddErrors(t *testing.T) {
	cases := []struct {
		name        string
		pokemonName string
		team        Team
		wantErr     bool
	}{
		{
			name:        "team is not full",
			pokemonName: "charmander",
			team:        NewTeam(3, 0, 0),
			wantErr:     false,
		},
		{
			name:        "team is almost full",
			pokemonName: "charmander",
			team:        NewTeam(1, 0, 0),
			wantErr:     false,
		},
		{
			name:        "team is full",
			pokemonName: "charmander",
			team:        NewTeam(0, 0, 0),
			wantErr:     true,
		},
		{
			name:        "pokemon with this name is already in the team",
			pokemonName: "duplicate name",
			team: Team{
				Pokemon: map[string]entities.Pokemon{
					"duplicate name": {},
				},
				MaxSize: 10,
			},
			wantErr: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotErr := c.team.Add(c.pokemonName, entities.Pokemon{})
			if gotErr != nil {
				if !c.wantErr {
					t.Errorf("Add() failed: %v", gotErr)
				}
				return
			}
			if c.wantErr {
				t.Fatal("Add() succeeded unexpectedly")
			}
		})
	}
}

func TestExperienceLoop(t *testing.T) {
	const maxTeamSize = 3
	cases := []struct {
		name               string
		passiveXPGain      uint
		passiveXPIntervall time.Duration
		waitDuration       time.Duration
		expected           map[string]int
	}{
		{
			name:               "team with 0 and non 0 values for current pokemon xp",
			passiveXPGain:      10,
			passiveXPIntervall: 10 * time.Millisecond,
			waitDuration:       59 * time.Millisecond,
			expected: map[string]int{
				"first":  50,
				"second": 150,
				"third":  50,
			},
		},
		{
			name:               "new team with 0 for intervall and passive xp gain",
			passiveXPGain:      0,
			passiveXPIntervall: 0,
			waitDuration:       20,
			expected: map[string]int{
				"first":  0,
				"second": 100,
				"third":  0,
			},
		},
		{
			name:               "new team with 0 for intervall and not 0 passive xp gain",
			passiveXPGain:      5,
			passiveXPIntervall: 0,
			waitDuration:       20,
			expected: map[string]int{
				"first":  0,
				"second": 100,
				"third":  0,
			},
		},
		{
			name:               "new team with 0 for passive xp gain and not 0 for intervall",
			passiveXPGain:      5,
			passiveXPIntervall: 0,
			waitDuration:       20,
			expected: map[string]int{
				"first":  0,
				"second": 100,
				"third":  0,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			team := NewTeam(maxTeamSize, c.passiveXPGain, c.passiveXPIntervall)
			team.Add("first", entities.Pokemon{CurrentExperience: 0})
			team.Add("second", entities.Pokemon{CurrentExperience: 100})
			team.Add("third", entities.Pokemon{CurrentExperience: 0})

			time.Sleep(c.waitDuration)

			for name, pokemon := range team.Pokemon {
				if c.expected[name] != pokemon.CurrentExperience {
					t.Errorf(
						"ExperienceLoop() mismatch for pokemon: %s\nwant: %v\ngot: %v\n",
						name,
						c.expected[name],
						pokemon.CurrentExperience,
					)
					return
				}
			}
		})
	}
}

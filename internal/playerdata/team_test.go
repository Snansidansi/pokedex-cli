package playerdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
			actual := NewTeam(c.maxsize)
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("NewTeam() mismatch (-want +got):\n%s", diff)
				return
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
			team:        NewTeam(3),
			wantErr:     false,
		},
		{
			name:        "team is almost full",
			pokemonName: "charmander",
			team:        NewTeam(1),
			wantErr:     false,
		},
		{
			name:        "team is full",
			pokemonName: "charmander",
			team:        NewTeam(0),
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

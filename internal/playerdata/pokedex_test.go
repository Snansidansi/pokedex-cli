package playerdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPokedex(t *testing.T) {
	cases := []struct {
		name            string
		keysToAdd       []string
		expectedPokedex Pokedex
	}{
		{
			name:      "Create empty pokedex",
			keysToAdd: []string{},
			expectedPokedex: Pokedex{
				Data: map[string]struct{}{},
			},
		},
		{
			name:      "Create pokedex with one init value",
			keysToAdd: []string{"pikachu"},
			expectedPokedex: Pokedex{
				Data: map[string]struct{}{
					"pikachu": {},
				},
			},
		},
		{
			name:      "Create pokedex with two init values",
			keysToAdd: []string{"charmander", "squirtle"},
			expectedPokedex: Pokedex{
				Data: map[string]struct{}{
					"charmander": {},
					"squirtle":   {},
				},
			},
		},
		{
			name:      "Create pokedex with duplicate init values",
			keysToAdd: []string{"charmander", "charmander"},
			expectedPokedex: Pokedex{
				Data: map[string]struct{}{
					"charmander": {},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pokedex := NewPokedex(c.keysToAdd...)
			if diff := cmp.Diff(c.expectedPokedex, pokedex); diff != "" {
				t.Errorf("NewPokedex() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func TestPokedex_Add(t *testing.T) {
	cases := []struct {
		name            string
		initialPokedex  Pokedex
		keyToAdd        string
		expectedPokedex Pokedex
	}{
		{
			name:            "Add new entry to empty pokedex",
			initialPokedex:  NewPokedex(),
			keyToAdd:        "pikachu",
			expectedPokedex: NewPokedex("pikachu"),
		},
		{
			name:            "Add new entry to non-empty pokedex",
			initialPokedex:  NewPokedex("charmander"),
			keyToAdd:        "bulbasaur",
			expectedPokedex: NewPokedex("charmander", "bulbasaur"),
		},
		{
			name:            "Add existing entry again (should not duplicate)",
			initialPokedex:  NewPokedex("squirtle"),
			keyToAdd:        "squirtle",
			expectedPokedex: NewPokedex("squirtle"),
		},
		{
			name:            "Add empty string as name",
			initialPokedex:  NewPokedex(),
			keyToAdd:        "",
			expectedPokedex: NewPokedex(""),
		},
		{
			name:            "Add name with special characters",
			initialPokedex:  NewPokedex(),
			keyToAdd:        "mr.mime",
			expectedPokedex: NewPokedex("mr.mime"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.initialPokedex.Add(c.keyToAdd)

			if diff := cmp.Diff(c.expectedPokedex, c.initialPokedex); diff != "" {
				t.Errorf("%s:\nAdd() mismatch (-want +got):\n%s", c.name, diff)
				return
			}
		})
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name       string
		pokedex    Pokedex
		checkValue string
		expected   bool
	}{
		{
			name:       "Empty pokedex does not contain value",
			pokedex:    NewPokedex(),
			checkValue: "pikachu",
			expected:   false,
		},
		{
			name:       "Pokedex contains value",
			pokedex:    NewPokedex("pikachu", "charmander"),
			checkValue: "charmander",
			expected:   true,
		},
		{
			name:       "Pokedex does not contain value",
			pokedex:    NewPokedex("pikachu", "charmander"),
			checkValue: "squirtle",
			expected:   false,
		},
	}

	for _, c := range cases {
		actual := c.pokedex.Contains(c.checkValue)

		if actual != c.expected {
			t.Errorf("Expected: %v\nActual: %v\n"+
				"For checkvalue: %s\nIn pokedex: %v\n",
				actual, c.expected, c.checkValue, c.pokedex)
			return
		}
	}
}

func TestGetAll(t *testing.T) {
	cases := []struct {
		name     string
		pokedex  Pokedex
		expected []string
	}{
		{
			name:     "Empty pokedex",
			pokedex:  NewPokedex(),
			expected: []string{},
		},
		{
			name:     "Not empty pokedex",
			pokedex:  NewPokedex("charmander", "pikachu", "squirtle"),
			expected: []string{"charmander", "pikachu", "squirtle"},
		},
		{
			name:     "Get all sorted",
			pokedex:  NewPokedex("squirtle", "pikachu", "charmander"),
			expected: []string{"charmander", "pikachu", "squirtle"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.pokedex.GetAll()

			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetAll() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		name     string
		pokedex  Pokedex
		expected bool
	}{
		{
			name:     "Empty pokedex",
			pokedex:  NewPokedex(),
			expected: true,
		},
		{
			name:     "Not empty pokedex",
			pokedex:  NewPokedex("pikachu", "charmander"),
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.pokedex.IsEmpty()

			if actual != c.expected {
				t.Errorf("Expected: %v\nActual: %v\nPokedex: %v\n", c.expected, actual, c.pokedex)
				return
			}
		})
	}
}

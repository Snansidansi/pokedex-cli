package pokeapi

import "testing"

func TestCalcCatchChance(t *testing.T) {
	pokemon := Pokemon{
		BaseExperience: 100,
	}
	cases := []struct {
		catchChanceMultiplier float64
		expected              int
	}{
		{
			catchChanceMultiplier: 1.0,
			expected:              43,
		},
		{
			catchChanceMultiplier: 1.5,
			expected:              62,
		},
		{
			catchChanceMultiplier: 2.0,
			expected:              72,
		},
		{
			catchChanceMultiplier: 5.0,
			expected:              89,
		},
		{
			catchChanceMultiplier: 255.0,
			expected:              100,
		},
	}

	for _, c := range cases {
		actualCatchChance := pokemon.CalcCatchChance(c.catchChanceMultiplier)
		if actualCatchChance != c.expected {
			t.Error("expected catchChance is different from actual")
			t.Errorf("Expected: %v\n", c.expected)
			t.Errorf("Actual: %v\n", actualCatchChance)
			return
		}
	}
}

func TestMasterBallMaxHasCatchChance(t *testing.T) {
	const expectedCatchChance = 100
	pokemons := []Pokemon{
		{
			BaseExperience: 0,
		},
		{
			BaseExperience: 1,
		},
		{
			BaseExperience: 50,
		},
		{
			BaseExperience: 100,
		},
		{
			BaseExperience: 1000,
		},
		{
			BaseExperience: 100000,
		},
		{
			BaseExperience: -1,
		},
	}

	for _, pokemon := range pokemons {
		masterBallCatchChanceMulti := GetPokeballs()["Master Ball"].CatchRateMultiplier
		actualCatchChance := pokemon.CalcCatchChance(masterBallCatchChanceMulti)

		if actualCatchChance != expectedCatchChance {
			t.Error("catch chance with Master Ball does not match the expected chance")
			t.Errorf("Expected: %v\n", expectedCatchChance)
			t.Errorf("Actual: %v\n", actualCatchChance)
			return
		}
	}
}

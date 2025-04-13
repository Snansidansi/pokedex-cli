package entities

import (
	"slices"
)

type PokeBall struct {
	Name                string
	Rarity              int
	MaxLootAmount       int
	CatchRateMultiplier float64
}

func GetPokeballs() map[string]PokeBall {
	return map[string]PokeBall{
		"Poké Ball": {
			Name:                "Poké Ball",
			Rarity:              12,
			MaxLootAmount:       5,
			CatchRateMultiplier: 1.0,
		},
		"Great Ball": {
			Name:                "Great Ball",
			Rarity:              5,
			MaxLootAmount:       2,
			CatchRateMultiplier: 1.5,
		},
		"Ultra Ball": {
			Name:                "Ultra Ball",
			Rarity:              3,
			MaxLootAmount:       1,
			CatchRateMultiplier: 2.0,
		},
		"Master Ball": {
			Name:                "Master Ball",
			Rarity:              1,
			MaxLootAmount:       1,
			CatchRateMultiplier: 255.0,
		},
	}
}

func GetPokeballsSorted() []PokeBall {
	pokeballsUnsorted := GetPokeballs()
	pokeballsSorted := make([]PokeBall, len(pokeballsUnsorted))

	i := 0
	for _, pokeball := range pokeballsUnsorted {
		pokeballsSorted[i] = pokeball
		i++
	}

	slices.SortFunc(pokeballsSorted, func(a, b PokeBall) int {
		return int(a.CatchRateMultiplier) - int(b.CatchRateMultiplier)
	})

	return pokeballsSorted
}

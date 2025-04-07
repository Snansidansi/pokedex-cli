package entities

type PokeBall struct {
	Name                string
	CatchRateMultiplier float64
}

func GetPokeballs() map[string]PokeBall {
	return map[string]PokeBall{
		"Poké Ball": {
			Name:                "Poké Ball",
			CatchRateMultiplier: 1.0,
		},
		"Great Ball": {
			Name:                "Great Ball",
			CatchRateMultiplier: 1.5,
		},
		"Ultra Ball": {
			Name:                "Ultra Ball",
			CatchRateMultiplier: 2.0,
		},
		"Master Ball": {
			Name:                "Master Ball",
			CatchRateMultiplier: 255.0,
		},
	}
}

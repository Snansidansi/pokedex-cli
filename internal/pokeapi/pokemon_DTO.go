package pokeapi

import "math/rand"

type PokemonDTO struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (p *PokemonDTO) Catch(pokeball PokeBall) (success bool) {
	if pokeball.Name == "Master Ball" {
		return true
	}

	catchChance := p.CalcCatchChance(pokeball.CatchRateMultiplier)

	randNum := rand.Intn(101) + 1
	catched := (randNum <= catchChance)

	return catched
}

func (p *PokemonDTO) CalcCatchChance(catchChanceMultiplier float64) int {
	if catchChanceMultiplier == GetPokeballs()["Master Ball"].CatchRateMultiplier {
		return 100
	}

	const minCatchChance = 5
	catchDifficulty := float64(p.BaseExperience) / (3.5 * catchChanceMultiplier * 0.5)

	catchChance := 100 - int(catchDifficulty)
	catchChance = max(minCatchChance, catchChance)

	return catchChance
}

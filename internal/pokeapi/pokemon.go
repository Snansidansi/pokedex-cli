package pokeapi

import "math/rand"

type Pokemon struct {
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
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
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

func (p *Pokemon) Catch(pokeball PokeBall) (success bool) {
	catchChance := p.CalcCatchChance(pokeball)

	randNum := rand.Intn(101) + 1
	catched := (randNum <= catchChance)

	return catched
}

func (p *Pokemon) CalcCatchChance(pokeball PokeBall) int {
	const minCatchChance = 5
	catchDifficulty := float64(p.BaseExperience) / (3.5 * (pokeball.CatchRateMultiplier * 0.5))

	catchChance := 100 - int(catchDifficulty)
	catchChance = max(minCatchChance, catchChance)

	return catchChance
}

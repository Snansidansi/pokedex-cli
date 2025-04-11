package entities

import (
	"fmt"
	"math"
	"strings"
)

type Stats struct {
	MaxHP     int
	CurrentHP int
	Damage    int
	Speed     int
}

type PokemonHP struct {
	Name      string
	CurrentHP int
	MaxHP     int
}

func SortPokemonHP(a, b PokemonHP) int {
	firstIsMaxHP := a.CurrentHP == a.MaxHP
	secondIsMaxHP := b.CurrentHP == b.MaxHP
	if firstIsMaxHP && !secondIsMaxHP {
		return 1
	} else if !firstIsMaxHP && secondIsMaxHP {
		return -1
	}

	hpDiff := a.CurrentHP - b.CurrentHP
	if hpDiff != 0 {
		return hpDiff
	}
	return strings.Compare(a.Name, b.Name)
}

type Pokemon struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	BaseExperience    int    `json:"base_experience"`
	CurrentExperience int    `json:"current_experience"`
	Height            int    `json:"height"`
	Order             int    `json:"order"`
	Weight            int    `json:"weight"`
	Abilities         []struct {
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
	Stats     Stats `json:"stats"`
	BaseStats Stats `json:"base_stats"`
	Types     []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	ImageUrl string `json:"official-artwork"`
}

func (pokemon Pokemon) Print() {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Current xp: %v\n", pokemon.CurrentExperience)
	fmt.Printf("Current level: %v\n", pokemon.GetLevel())
	fmt.Printf(" -> xp to next level (%v): %v\n", pokemon.GetLevel()+1, pokemon.GetXPForNextLevel())
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Abilies:")
	for i := range pokemon.Abilities {
		fmt.Printf(" - %s\n", pokemon.Abilities[i].Ability.Name)
	}

	fmt.Println("Forms:")
	for i := range pokemon.Forms {
		fmt.Printf(" - %s\n", pokemon.Forms[i].Name)
	}

	fmt.Println("Stats:")
	fmt.Printf(" - Max hp: %v\n", pokemon.Stats.MaxHP)
	fmt.Printf(" - Current hp: %v\n", pokemon.Stats.CurrentHP)
	fmt.Printf(" - Damage: %v\n", pokemon.Stats.Damage)
	fmt.Printf(" - Speed: %v\n", pokemon.Stats.Speed)

	fmt.Println("Types:")
	for i := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon.Types[i].Type.Name)
	}
}

func (pokemon *Pokemon) AddExperience(amount int) {
	pokemon.CurrentExperience += amount
	pokemon.RecalculateStats()
}

const XPForLevelTuningFactor = 0.4

func (pokemon Pokemon) GetLevel() int {
	level := XPForLevelTuningFactor * math.Sqrt(float64(pokemon.CurrentExperience))
	return int(math.Round(level))
}

func (pokemon *Pokemon) SetLevel(level int) {
	xp := math.Ceil(math.Pow(float64(level)/XPForLevelTuningFactor, 2))
	pokemon.CurrentExperience = int(xp)
}

func (pokemon Pokemon) GetXPForNextLevel() int {
	currentLevel := pokemon.GetLevel()
	nextLevel := currentLevel + 1

	totalXpForNextLevel := math.Pow(float64(nextLevel)/XPForLevelTuningFactor, 2)

	neededXp := int(math.Ceil(totalXpForNextLevel)) - pokemon.CurrentExperience
	return neededXp
}

func (pokemon *Pokemon) RecalculateStats() {
	currentLevel := pokemon.GetLevel()
	oldMaxHP := pokemon.Stats.MaxHP

	pokemon.Stats.MaxHP = pokemon.BaseStats.MaxHP + int(math.Pow(float64(currentLevel), 1.5))
	pokemon.Stats.Damage = pokemon.BaseStats.Damage + int(math.Pow(float64(currentLevel), 1.2))

	HPUpgrade := pokemon.Stats.MaxHP - oldMaxHP
	pokemon.Stats.CurrentHP += HPUpgrade
}

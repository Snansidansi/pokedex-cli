package entities

import (
	"fmt"
	"math"
)

type Stats struct {
	MaxHP     int
	CurrentHP int
	Damage    int
	Speed     int
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

const XPForLevelTuningFactor = 0.4

func (pokemon Pokemon) GetLevel() int {
	level := XPForLevelTuningFactor * math.Sqrt(float64(pokemon.CurrentExperience))
	return int(math.Round(level))
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
	isFullLive := pokemon.Stats.MaxHP == pokemon.Stats.CurrentHP

	pokemon.Stats.MaxHP = pokemon.BaseStats.MaxHP + int(math.Pow(float64(currentLevel), 1.5))
	if isFullLive {
		pokemon.Stats.CurrentHP = pokemon.Stats.MaxHP
	}

	pokemon.Stats.Damage = pokemon.BaseStats.Damage + int(math.Pow(float64(currentLevel), 1.7))
}

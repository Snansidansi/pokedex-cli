package entities

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type Stats struct {
	MaxHP     int
	CurrentHP int
	Damage    int
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
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	BaseExperience    int      `json:"base_experience"`
	CurrentExperience int      `json:"current_experience"`
	Height            int      `json:"height"`
	Weight            int      `json:"weight"`
	Stats             Stats    `json:"stats"`
	BaseStats         Stats    `json:"base_stats"`
	Types             []string `json:"types"`
	ImageUrl          string   `json:"official-artwork"`
}

func (p Pokemon) Catch(pokeball PokeBall) (success bool) {
	if pokeball.Name == "Master Ball" {
		return true
	}

	catchChance := p.CalcCatchChance(pokeball.CatchRateMultiplier)

	randNum := rand.Intn(100) + 1
	catched := (randNum <= catchChance)

	return catched
}

func (p Pokemon) CalcCatchChance(catchChanceMultiplier float64) int {
	if catchChanceMultiplier == GetPokeballs()["Master Ball"].CatchRateMultiplier {
		return 100
	}

	const minCatchChance = 5
	catchDifficulty := float64(p.BaseExperience) / (3.5 * catchChanceMultiplier * 0.5)

	catchChance := 100 - int(catchDifficulty)
	catchChance = max(minCatchChance, catchChance)

	return catchChance
}

func (pokemon Pokemon) Print() {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Current xp: %v\n", pokemon.CurrentExperience)
	fmt.Printf("Current level: %v\n", pokemon.GetLevel())
	fmt.Printf(" -> xp to next level (%v): %v\n", pokemon.GetLevel()+1, pokemon.GetXPForNextLevel())
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	fmt.Printf(" - Max hp: %v\n", pokemon.Stats.MaxHP)
	fmt.Printf(" - Current hp: %v\n", pokemon.Stats.CurrentHP)
	fmt.Printf(" - Damage: %v\n", pokemon.Stats.Damage)

	fmt.Println("Types:")
	for _, name := range pokemon.Types {
		fmt.Printf(" - %s\n", name)
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
	pokemon.RecalculateStats()
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

func (pokemon *Pokemon) TakeDamage(amount int) (pokemonDied bool) {
	pokemon.Stats.CurrentHP -= amount
	if pokemon.Stats.CurrentHP < 0 {
		return true
	}
	return false
}

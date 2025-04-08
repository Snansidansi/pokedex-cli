package entities

import "fmt"

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

func (pokemon Pokemon) Print() {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Current xp: %v\n", pokemon.CurrentExperience)
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
	for i := range pokemon.Stats {
		stat := pokemon.Stats[i]
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for i := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon.Types[i].Type.Name)
	}
}

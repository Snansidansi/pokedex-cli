package playerdata

import (
	"math/rand"

	"github.com/snansidansi/pokedex-cli/internal/entities"
)

type PokeballInv map[string]int

type pokeballAmountMapping struct {
	Name   string
	Amount int
}

func NewPokeballInv() PokeballInv {
	newPokeballInv := PokeballInv{}
	for name := range entities.GetPokeballs() {
		newPokeballInv[name] = 0
	}
	return newPokeballInv
}

func (pokeballInv PokeballInv) IsEmpty() bool {
	for _, amount := range pokeballInv {
		if amount > 0 {
			return false
		}
	}

	return true
}

func (PokeballInv PokeballInv) FindPokeballLoot() []pokeballAmountMapping {
	loot := []pokeballAmountMapping{}
	for _, pokeball := range entities.GetPokeballs() {
		randFind := rand.Intn(100) + 1
		if randFind > pokeball.Rarity {
			continue
		}

		randAmount := rand.Intn(pokeball.MaxLootAmount) + 1
		PokeballInv[pokeball.Name] += randAmount

		lootEntry := pokeballAmountMapping{
			Name:   pokeball.Name,
			Amount: randAmount,
		}
		loot = append(loot, lootEntry)
	}

	return loot
}

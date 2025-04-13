package playerdata

type PokeballInv map[string]int

func (pokeballInv PokeballInv) IsEmpty() bool {
	for _, amount := range pokeballInv {
		if amount > 0 {
			return false
		}
	}

	return true
}

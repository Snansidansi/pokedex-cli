package playerdata

import "slices"

type Pokedex struct {
	Data map[string]struct{} `json:"data"`
}

func NewPokedex(initData ...string) *Pokedex {
	pokedex := &Pokedex{
		Data: map[string]struct{}{},
	}

	for _, data := range initData {
		pokedex.Add(data)
	}

	return pokedex
}

func (p *Pokedex) Add(name string) {
	p.Data[name] = struct{}{}
}

func (p *Pokedex) Contains(name string) bool {
	_, ok := p.Data[name]
	return ok
}

func (p *Pokedex) GetAll() []string {
	pokemonNames := make([]string, len(p.Data))
	i := 0
	for name := range p.Data {
		pokemonNames[i] = name
		i++
	}

	slices.Sort(pokemonNames)

	return pokemonNames
}

func (p *Pokedex) IsEmpty() bool {
	return len(p.Data) == 0
}

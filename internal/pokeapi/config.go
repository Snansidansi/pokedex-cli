package pokeapi

type Config struct {
	Client          Client
	NextLocationURL *string
	PrevLocationURL *string
	Pokedex         map[string]Pokemon
}

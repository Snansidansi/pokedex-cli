package pokeapi

type Config struct {
	Client          Client
	NextLocationURL *string
	PrevLocationURL *string
	PlayerData      PlayerData
	SaveFile        SaveFile
}

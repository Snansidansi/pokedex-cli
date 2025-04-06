package playerdata

type PlayerData struct {
	Pokedex *Pokedex `json:"pokedex"`
}

func NewPlayerData() PlayerData {
	return PlayerData{
		Pokedex: NewPokedex(),
	}
}

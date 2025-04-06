package playerdata

type PlayerData struct {
	Pokedex *Pokedex `json:"pokedex"`
	Pokebox *Pokebox `json:"pokebox"`
}

func NewPlayerData() PlayerData {
	return PlayerData{
		Pokedex: NewPokedex(),
		Pokebox: &Pokebox{},
	}
}

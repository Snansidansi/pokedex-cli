package playerdata

type PlayerData struct {
	Pokedex Pokedex `json:"pokedex"`
	Pokebox Pokebox `json:"pokebox"`
	Team    Team    `json:"team"`
}

func NewPlayerData() PlayerData {
	return PlayerData{
		Pokedex: NewPokedex(),
		Pokebox: Pokebox{},
		Team:    Team{},
	}
}

package playerdata

type PlayerData struct {
	Pokedex Pokedex `json:"pokedex"`
	Pokebox Pokebox `json:"pokebox"`
	Team    Team    `json:"team"`
}

const (
	maxTeamSize        = 6
	passiveXPGain      = 0
	passiveXPIntervall = 0
)

func NewPlayerData() PlayerData {
	return PlayerData{
		Pokedex: NewPokedex(),
		Pokebox: Pokebox{},
		Team:    NewTeam(maxTeamSize, passiveXPGain, passiveXPIntervall),
	}
}

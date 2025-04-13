package playerdata

import "time"

type PlayerData struct {
	Pokedex     Pokedex     `json:"pokedex"`
	Pokebox     Pokebox     `json:"pokebox"`
	Team        Team        `json:"team"`
	PokeballInv PokeballInv `json:"pokeball_inventory"`
}

const (
	maxTeamSize        = 6
	passiveXPGain      = 1
	passiveXPIntervall = 2 * time.Second
)

func NewPlayerData() PlayerData {
	return PlayerData{
		Pokedex:     NewPokedex(),
		Pokebox:     Pokebox{},
		Team:        NewTeam(maxTeamSize, passiveXPGain, passiveXPIntervall),
		PokeballInv: NewPokeballInv(),
	}
}

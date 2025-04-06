package pokeapi

import "github.com/snansidansi/pokedex-cli/internal/playerdata"

type Config struct {
	Client          Client
	NextLocationURL *string
	PrevLocationURL *string
	PlayerData      playerdata.PlayerData
	SaveFile        SaveFile
}

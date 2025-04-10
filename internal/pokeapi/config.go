package pokeapi

import "github.com/snansidansi/pokedex-cli/internal/playerdata"

type Config struct {
	Client          Client
	NextLocationURL *string
	PrevLocationURL *string
	CurrentLocation int
	PlayerData      playerdata.PlayerData
	SaveFile        SaveFile
}

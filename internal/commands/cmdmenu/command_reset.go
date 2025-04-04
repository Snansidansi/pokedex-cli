package cmdmenu

import (
	"errors"
	"fmt"
	"os"

	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
)

func commandReset(conf *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("type 'reset confirm' to reset your progress")
	}

	confirmation := args[0]
	if confirmation != "confirm" {
		return errors.New("type 'reset confirm' to reset your progress")
	}

	saveFilePath := conf.SaveFile.Dir + "/" + conf.SaveFile.Name
	if err := os.Remove(saveFilePath); err != nil {
		return err
	}

	conf.Pokedex = make(map[string]pokeapi.Pokemon)
	fmt.Println("resetted data successfuly")

	return nil
}

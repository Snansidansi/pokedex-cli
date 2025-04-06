package pokeapi

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/snansidansi/pokedex-cli/internal/playerdata"
)

type SaveFile struct {
	Dir  string
	Name string
}

func (c *Config) Save() error {
	jsonData, err := json.Marshal(c.PlayerData)
	if err != nil {
		return err
	}

	tempFilePath := c.SaveFile.Dir + "/tmp-savefile.txt"
	if err := os.WriteFile(tempFilePath, jsonData, 0644); err != nil {
		os.Remove(tempFilePath)
		return err
	}

	saveFilePath := c.SaveFile.Dir + "/" + c.SaveFile.Name
	os.Remove(saveFilePath)
	os.Rename(tempFilePath, saveFilePath)

	return nil
}

func (c *Config) Load() error {
	saveFilePath := c.SaveFile.Dir + "/" + c.SaveFile.Name

	data, err := os.ReadFile(saveFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	loadedData := playerdata.PlayerData{}
	if err := json.Unmarshal(data, &loadedData); err != nil {
		return err
	}

	c.PlayerData = loadedData

	return nil
}

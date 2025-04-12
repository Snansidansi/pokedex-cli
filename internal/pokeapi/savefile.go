package pokeapi

import (
	"encoding/json"
	"errors"
	"os"
	"path"

	"github.com/snansidansi/pokedex-cli/internal/playerdata"
)

const saveFileName = ".pokedex-cli-saveFile.json"

type SaveFile struct {
	Dir string
}

func (c *Config) Save() error {
	c.PlayerData.Team.Mu.Lock()
	jsonData, err := json.Marshal(c.PlayerData)
	if err != nil {
		return err
	}
	c.PlayerData.Team.Mu.Unlock()

	tempFilePath := c.SaveFile.Dir + "/.tmp-savefile.json"
	if err := os.WriteFile(tempFilePath, jsonData, 0644); err != nil {
		os.Remove(tempFilePath)
		return err
	}

	saveFilePath := c.getSaveFilePath()
	os.Remove(saveFilePath)
	os.Rename(tempFilePath, saveFilePath)

	return nil
}

func (c *Config) Load() error {
	saveFilePath := c.getSaveFilePath()

	data, err := os.ReadFile(saveFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	loadedData := playerdata.NewPlayerData()
	loadedData.Team.Mu.Lock()
	if err := json.Unmarshal(data, &loadedData); err != nil {
		return err
	}
	c.PlayerData = loadedData
	loadedData.Team.Mu.Unlock()

	return nil
}

func (c *Config) Reset() error {
	saveFilePath := c.getSaveFilePath()
	if err := os.Remove(saveFilePath); err != nil {
		return err
	}

	c.PlayerData = playerdata.NewPlayerData()
	return nil
}

func (c *Config) getSaveFilePath() string {
	return path.Join(c.SaveFile.Dir, saveFileName)
}

package pokeapi

import (
	"encoding/json"
	"errors"
	"os"
)

type SaveFile struct {
	Dir  string
	Name string
}

type SaveData struct {
	Pokedex map[string]Pokemon `json:"pokedex"`
}

func (c *Config) Save() error {
	data := SaveData{
		c.Pokedex,
	}

	jsonData, err := json.Marshal(data)
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

	saveData := SaveData{}
	if err := json.Unmarshal(data, &saveData); err != nil {
		return err
	}

	c.Pokedex = saveData.Pokedex

	return nil
}

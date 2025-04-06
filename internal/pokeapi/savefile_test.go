package pokeapi

import (
	"reflect"
	"testing"

	"github.com/snansidansi/pokedex-cli/internal/playerdata"
)

func TestSaveAndLoadData(t *testing.T) {
	tempDir := t.TempDir()
	expectedPlayerData := playerdata.NewPlayerData()
	expectedPlayerData.Pokedex.Add("testmon1")
	expectedPlayerData.Pokedex.Add("testmon2")

	saveFile := SaveFile{
		Dir:  tempDir,
		Name: "testSaveFile.txt",
	}

	inputConfig := Config{
		PlayerData: expectedPlayerData,
		SaveFile:   saveFile,
	}

	outputConfig := Config{
		SaveFile: saveFile,
	}

	if err := inputConfig.Save(); err != nil {
		t.Error("error occured while saving config")
		t.Error(err)
		return
	}

	if err := outputConfig.Load(); err != nil {
		t.Error("error occured while loading config")
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(inputConfig.PlayerData, outputConfig.PlayerData) {
		t.Error("saved data and loaded data are unequal")
		t.Errorf("saved data: %v", inputConfig.PlayerData)
		t.Errorf("loaded data: %v", outputConfig.PlayerData)
		return
	}
}

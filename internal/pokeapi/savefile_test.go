package pokeapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/snansidansi/pokedex-cli/internal/playerdata"
)

func TestSaveAndLoadData(t *testing.T) {
	tempDir := t.TempDir()
	expectedPlayerData := playerdata.NewPlayerData()
	expectedPlayerData.Pokedex.Add("testmon1")
	expectedPlayerData.Pokedex.Add("testmon2")

	saveFile := SaveFile{
		Dir: tempDir,
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

	if diff := cmp.Diff(
		inputConfig.PlayerData,
		outputConfig.PlayerData,
		cmpopts.IgnoreFields(playerdata.Team{}, "Mu"),
	); diff != "" {
		t.Errorf("saved and loaded data are unequal (-want +got):\n%s", diff)
		return
	}
}

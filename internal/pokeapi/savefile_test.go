package pokeapi

import (
	"reflect"
	"testing"
)

func TestSaveAndLoadData(t *testing.T) {
	tempDir := t.TempDir()
	expectedPokedex := map[string]Pokemon{
		"testmon": {
			Name:           "testmon",
			Weight:         23,
			Height:         35,
			BaseExperience: 67,
			IsDefault:      true,
		},
		"testmon2": {

			Name:           "testmon2",
			Weight:         65,
			Height:         96,
			BaseExperience: 111,
			IsDefault:      false,
		},
	}
	saveFile := SaveFile{
		Dir:  tempDir,
		Name: "testSaveFile.txt",
	}

	inputConfig := Config{
		Pokedex:  expectedPokedex,
		SaveFile: saveFile,
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

	if !reflect.DeepEqual(inputConfig.Pokedex, outputConfig.Pokedex) {
		t.Error("saved data and loaded data are unequal")
		t.Errorf("saved data: %v", inputConfig.Pokedex)
		t.Errorf("loaded data: %v", outputConfig.Pokedex)
		return
	}
}

package playerdata

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/entities"
)

type Team struct {
	Pokemon map[string]entities.Pokemon `json:"pokemon"`
	Mu      *sync.Mutex                 `json:"-"`
	MaxSize uint                        `json:"max_size"`
}

func NewTeam(maxSize uint, passiveXPGain uint, passiveXPIntervall time.Duration) Team {
	team := Team{
		Pokemon: map[string]entities.Pokemon{},
		Mu:      &sync.Mutex{},
		MaxSize: maxSize,
	}

	if passiveXPGain > 0 && passiveXPIntervall > 0 {
		go team.experienceLoop(passiveXPIntervall, passiveXPGain)
	}

	return team
}

func (team Team) Add(pokemonName string, pokemon entities.Pokemon) error {
	if len(team.Pokemon) >= int(team.MaxSize) {
		return TeamIsFullError{
			Message: fmt.Sprintf("team is full (max size: %v)\n", team.MaxSize),
		}
	}

	if _, ok := team.Pokemon[pokemonName]; ok {
		return fmt.Errorf("pokemon with that name already exists in team: %s", pokemonName)
	}

	team.Pokemon[pokemonName] = pokemon

	return nil
}

func (team Team) experienceLoop(intervall time.Duration, xpGain uint) {
	ticker := time.NewTicker(intervall)
	defer ticker.Stop()

	for {
		<-ticker.C

		team.Mu.Lock()
		for name, pokemon := range team.Pokemon {
			pokemon.CurrentExperience += int(xpGain)
			team.Pokemon[name] = pokemon
		}
		team.Mu.Unlock()
	}
}

func (team Team) Get(pokemonName string) (pokemon entities.Pokemon, ok bool) {
	team.Mu.Lock()
	pokemon, ok = team.Pokemon[pokemonName]
	team.Mu.Unlock()
	return
}

func (team Team) GetAllNamesSorted() []string {
	team.Mu.Lock()
	defer team.Mu.Unlock()

	pokemonNames := make([]string, len(team.Pokemon))
	i := 0
	for name := range team.Pokemon {
		pokemonNames[i] = name
		i++
	}

	slices.Sort(pokemonNames)
	return pokemonNames
}

func (team Team) Update(pokemonName string, newPokemon entities.Pokemon) error {
	team.Mu.Lock()
	defer team.Mu.Unlock()

	if _, ok := team.Pokemon[pokemonName]; !ok {
		return errors.New("pokemon cannot be updated: does not exist")
	}

	team.Pokemon[pokemonName] = newPokemon

	return nil
}

func (team Team) Rename(oldName, newName string) error {
	team.Mu.Lock()
	defer team.Mu.Unlock()

	pokemon, ok := team.Pokemon[oldName]
	if !ok {
		return errors.New("pokemon cannot be renamed: does not exist")
	}

	if _, ok := team.Pokemon[newName]; ok {
		return errors.New("pokemon cannot be renamed: new name already exists")
	}

	team.Pokemon[newName] = pokemon
	delete(team.Pokemon, oldName)

	return nil
}

func (team Team) Delete(pokemonName string) {
	team.Mu.Lock()
	delete(team.Pokemon, pokemonName)
	team.Mu.Unlock()
}

func (team Team) Size() int {
	team.Mu.Lock()
	defer team.Mu.Unlock()
	return len(team.Pokemon)
}

package cmdfight

import (
	"errors"
	"fmt"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/entities"
	"github.com/snansidansi/pokedex-cli/internal/playerdata"
	"github.com/snansidansi/pokedex-cli/internal/pokeapi"
	"github.com/snansidansi/pokedex-cli/internal/repl"
)

const timeBetweenFightMessages = 333 * time.Millisecond

func commandAttack(conf *pokeapi.Config, _ ...string) error {
	team := &conf.PlayerData.Team
	if team.ActivePokemon == nil {
		return errors.New("You do not have an active pokemon. Select one with the select command.")
	}

	enemyPokemon := team.CurrentEnemy
	teamPokemon, ok := team.Get(*team.ActivePokemon)
	if !ok {
		return errors.New("unexpected error: your active pokemon is not in your team anymore")
	}

	if teamPokemon.Stats.CurrentHP <= 0 {
		return errors.New("Your selected pokemon was already defeated. Select a different one.")
	}

	err := teamPokemonAttack(enemyPokemon, &teamPokemon, team)
	if err != nil {
		return err
	}

	time.Sleep(timeBetweenFightMessages)

	err = enemyPokemonAttack(enemyPokemon, team)
	if err != nil {
		return err
	}

	return nil
}

func teamPokemonAttack(enemyPokemon, teamPokemon *entities.Pokemon, team *playerdata.Team) error {
	enemyDied := enemyPokemon.TakeDamage(teamPokemon.Stats.Damage)
	fmt.Printf("%s attacks %s\n", *team.ActivePokemon, enemyPokemon.Name)
	time.Sleep(timeBetweenFightMessages)
	fmt.Printf("%s takes %v damage\n", enemyPokemon.Name, teamPokemon.Stats.Damage)
	time.Sleep(timeBetweenFightMessages)

	if enemyDied {
		fmt.Println("")
		fmt.Printf("%s was defeated by %s!\n", enemyPokemon.Name, *team.ActivePokemon)

		xpGain := enemyPokemon.BaseExperience * enemyPokemon.GetLevel()
		xpGain = int(float64(xpGain) * 0.26)

		fmt.Printf("You won the fight! All the Pokemon in your team gain %vxp.\n", xpGain)
		team.AddExperience(xpGain)
		team.WonFight = true
		return repl.ExitReplError{}
	}

	fmt.Printf("Hp of %s: %v / %vhp\n", enemyPokemon.Name, enemyPokemon.Stats.CurrentHP, enemyPokemon.Stats.MaxHP)
	fmt.Println("")
	return nil
}

func enemyPokemonAttack(enemyPokemon *entities.Pokemon, team *playerdata.Team) error {
	teamPokemonDied, err := team.DamagePokemon(*team.ActivePokemon, enemyPokemon.Stats.Damage)
	if err != nil {
		return err
	}

	fmt.Printf("%s attacks %s\n", enemyPokemon.Name, *team.ActivePokemon)
	time.Sleep(timeBetweenFightMessages)
	fmt.Printf("%s takes %v damage\n", *team.ActivePokemon, enemyPokemon.Stats.Damage)
	time.Sleep(timeBetweenFightMessages)

	if teamPokemonDied {
		fmt.Println("")
		fmt.Printf("%s was defeated.\n", *team.ActivePokemon)

		if !team.HasAliveMembers() {
			fmt.Println("")
			fmt.Println("All the pokemon in your team are defeated.")
			fmt.Println("You lost the fight, you should heal your pokemon.")
			team.WonFight = false
			return repl.ExitReplError{}
		}

		fmt.Println("Select a new pokemon from you team to fight")
		return nil
	}

	teamPokemon, _ := team.Get(*team.ActivePokemon)
	fmt.Printf("Hp of %s: %v / %vhp\n", *team.ActivePokemon, teamPokemon.Stats.CurrentHP, teamPokemon.Stats.MaxHP)
	return nil
}

//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{
		name: "Mario",
		health: 0,
		maxHealth: 100,
		energy: 0,
		maxEnergy: 100,
	}
}


func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(120)

	if player.health > player.maxHealth {
		t.Fatalf("Player health %v can not exceed max health of %v", player.health, player.maxHealth)
	}

	player.applyDamage(player.maxHealth + 1)

	if player.health < 0 {
		t.Fatalf("Player health %v can not be less then zero", player.health)
	}

	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Max health: %v", player.health, player.maxHealth)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.addEnergy(120)

	if player.energy > player.maxEnergy {
		t.Fatalf("Player energy %v can not exceed max energy of %v", player.energy, player.maxEnergy)
	}

	player.consumeEnergy(player.energy + 1)

	if player.energy < 0 {
		t.Fatalf("Player energy %v can not be less then zero", player.energy)
	}

	if player.energy > player.maxEnergy {
		t.Fatalf("Energy: %v. Max energy: %v", player.energy, player.maxEnergy)
	}
}
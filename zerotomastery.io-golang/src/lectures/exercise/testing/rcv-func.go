package main

import "fmt"

type Player struct {
	name string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount

	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}

	fmt.Println(player.name, "Add", amount, "health ->", player.health)
}

func (player *Player) applyDamage(amount uint) {
	if player.health - amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}

	fmt.Println(player.name, "DAMAGE", amount, "->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount

	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}

	fmt.Println(player.name, "Add", amount, "energy ->", player.energy)
}

func (player *Player) consumeEnergy(amount uint) {
	if player.energy - amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}

	fmt.Println(player.name, "consumed energy", amount, "->", player.energy)
}

func main() {
	player := Player{
		name: "Mario",
		health: 0,
		maxHealth: 100,
		energy: 0,
		maxEnergy: 100,
	}

	player.addHealth(50)
	player.applyDamage(25)
	player.addEnergy(50)
	player.consumeEnergy(20)
	player.consumeEnergy(200)
}

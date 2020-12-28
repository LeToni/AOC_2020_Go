package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Attacker struct {
	hp, rawDamage, armor int
}

type Item struct {
	cost, damage, armor int
}

var (
	Weapons = []*Item{
		{cost: 8, damage: 4, armor: 0},
		{cost: 10, damage: 5, armor: 0},
		{cost: 25, damage: 6, armor: 0},
		{cost: 40, damage: 7, armor: 0},
		{cost: 74, damage: 8, armor: 0},
	}
	Armor = []*Item{
		{cost: 13, damage: 0, armor: 1},
		{cost: 31, damage: 0, armor: 2},
		{cost: 53, damage: 0, armor: 3},
		{cost: 75, damage: 0, armor: 4},
		{cost: 102, damage: 0, armor: 5},
		{cost: 0, damage: 0, armor: 0},
	}
	Rings = []*Item{
		{cost: 25, damage: 1, armor: 0},
		{cost: 50, damage: 2, armor: 0},
		{cost: 100, damage: 3, armor: 0},
		{cost: 20, damage: 0, armor: 1},
		{cost: 40, damage: 0, armor: 2},
		{cost: 80, damage: 0, armor: 3},
		{cost: 0, damage: 0, armor: 0},
		{cost: 0, damage: 0, armor: 0},
	}

	Boss   Attacker
	Player Attacker
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			enemyHP, enemyDamage, enemyArmor int
		)
		input := scanner.Text()
		if n, _ := fmt.Sscanf(input, "Hit Points: %d", &enemyHP); n == 1 {
			Boss.hp = enemyHP
		} else if n, _ := fmt.Sscanf(input, "Damage: %d", &enemyDamage); n == 1 {
			Boss.rawDamage = enemyDamage
		} else if n, _ := fmt.Sscanf(input, "Armor: %d", &enemyArmor); n == 1 {
			Boss.armor = enemyArmor
		}
	}
}

func WinAgainstBoss() bool {
	heroDamage := Player.rawDamage - Boss.armor
	bossDamage := Boss.rawDamage - Player.armor

	if heroDamage < 1 {
		heroDamage = 1
	}
	if bossDamage < 1 {
		bossDamage = 1
	}

	roundsForPlayerWin := math.Ceil(float64(Boss.hp) / float64(heroDamage))
	roundsForBossWin := math.Ceil(float64(Player.hp) / float64(bossDamage))

	return roundsForPlayerWin <= roundsForBossWin
}

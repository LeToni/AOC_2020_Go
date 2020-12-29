package main

type Player struct {
	hp, mana, armor int
}

type Boss struct {
	hp, damage int
}

type Spell struct {
	name   string
	mana   int
	damage int
	heal   int
	effect string
}

type Effect struct {
	Duration, Remaining int
	Active              bool
}

var (
	spells = []*Spell{
		{name: "Magic Missole", mana: 53, damage: 4, heal: 0, effect: "None"},
		{name: "Drain", mana: 73, damage: 2, heal: 2, effect: "None"},
		{name: "Shield", mana: 113, damage: 0, heal: 0, effect: "Shield"},
		{name: "Poison", mana: 173, damage: 3, heal: 0, effect: "Poison"},
		{name: "Recharge", mana: 229, damage: 0, heal: 0, effect: "Recharge"},
	}

	player             Player
	hpPlayer           int = 50
	manaPlayer         int = 500
	boss               Boss
	hpBoss, bossDamage int
)

func main() {

}

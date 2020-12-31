package main

type Wizard struct {
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
	duration, remaining int
	active              bool
}

func (wizard *Wizard) isDead() bool {
	// wizard can be also considered dead if he has not enough mana to cast spell
	if wizard.hp <= 0 || wizard.mana < 53 {
		return true
	} else {
		return false
	}
}

func (wizard *Wizard) canCast(spell *Spell) bool {
	if spell.mana > wizard.mana {
		return false
	}

	if spell.name == "Shield" || spell.name == "Poison" || spell.name == "Recharge" {
		return !effects[spell.name].active
	}

	return true
}

func (boss *Boss) isDead() bool {
	if boss.hp <= 0 {
		return true
	} else {
		return false
	}
}

func (boss *Boss) attacks(wizard *Wizard) {
	if boss.damage-wizard.armor < 1 {
		wizard.hp = wizard.hp - 1
	} else {
		wizard.hp = wizard.hp - boss.damage
	}
}

func EffectTimer() {
	for _, effect := range effects {
		if effect.active && effect.remaining > 0 {
			effect.remaining--
		} else {
			effect.remaining = effect.duration
			effect.active = false
		}
	}
}

var (
	spells = []*Spell{
		{name: "Magic Missole", mana: 53, damage: 4, heal: 0, effect: "None"},
		{name: "Drain", mana: 73, damage: 2, heal: 2, effect: "None"},
		{name: "Shield", mana: 113, damage: 0, heal: 0, effect: "Shield"},
		{name: "Poison", mana: 173, damage: 3, heal: 0, effect: "Poison"},
		{name: "Recharge", mana: 229, damage: 0, heal: 0, effect: "Recharge"},
	}

	effects = map[string]*Effect{
		"Shield":   {duration: 6, remaining: 6, active: false},
		"Poison":   {duration: 6, remaining: 6, active: false},
		"Recharge": {duration: 5, remaining: 5, active: false},
	}

	wizard Wizard
	boss   Boss
)

func main() {

}

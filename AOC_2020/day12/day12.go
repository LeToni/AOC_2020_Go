package main

import "fmt"

const (
	N int = iota
	E
	S
	W
	R
	L
	F
)

type Instruction struct {
	action int
	value  int
}

type Ship struct {
	facing int
	x, y   int
}

func (ship *Ship) navigate(cardinalDirection rune, units int) {
	switch cardinalDirection {
	case 'N':
		ship.y = ship.y + units
	case 'E':
		ship.x = ship.x + units
	case 'S':
		ship.y = ship.y - units
	case 'W':
		ship.x = ship.x - units
	default:
		err := fmt.Errorf("Not a valid cardinal direction: %d", cardinalDirection)
		panic(err)
	}
}

func (ship *Ship) move(units int) {
	switch ship.facing {
	case E:
		ship.x = ship.x + units
	case S:
		ship.y = ship.y - units
	case W:
		ship.x = ship.x - units
	default:
		ship.y = ship.y + units
	}
}

func (ship *Ship) turn(direction rune) {
	if direction == 'R' {
		ship.facing = (ship.facing + 1) % 4
	} else {
		ship.facing = ((ship.facing - 1) + 4) % 4
	}
}

func main() {

}

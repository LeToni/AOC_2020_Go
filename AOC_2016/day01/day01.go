package main

const (
	N int = 0
	E int = 1
	S int = 2
	W int = 3
)

type Passenger struct {
	x, y   int
	facing int
}

func (passenger *Passenger) Walk(direction string, blocks int) {

	passenger.Turn(direction)

	switch passenger.facing {
	case E:
		passenger.x = passenger.x + blocks
	case S:
		passenger.y = passenger.y - blocks
	case W:
		passenger.x = passenger.x - blocks
	default:
		passenger.y = passenger.y + blocks
	}
}

func (passenger *Passenger) Turn(direction string) {

	if direction == "R" {
		passenger.facing = (passenger.facing + 1) % 4
	} else {
		passenger.facing = ((passenger.facing - 1) + 4) % 4
	}
}

func main() {

}

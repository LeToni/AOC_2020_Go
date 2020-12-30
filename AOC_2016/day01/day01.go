package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func distanceToHQ(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	passenger := Passenger{x: 0, y: 0, facing: N}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		directions := strings.Split(input, ",")

		filterBlocks := regexp.MustCompile(`(\d+)`)
		for _, direction := range directions {
			blocks, _ := strconv.Atoi(filterBlocks.FindAllString(direction, -1)[0])

			if strings.Contains(direction, "R") {
				passenger.Walk("R", blocks)
			} else {
				passenger.Walk("L", blocks)
			}
		}
	}

	distance := distanceToHQ(passenger.x, passenger.y)
	fmt.Println("Distance from landing zone to HQ: ", distance)
}

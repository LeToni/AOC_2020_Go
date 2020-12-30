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

func (passenger *Passenger) Walk(blocks int) {
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

type Location struct {
	x, y int
}

func distanceToHQ(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

var (
	visited = make(map[Location]bool)
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	passenger := Passenger{x: 0, y: 0, facing: N}
	secondSolved := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		directions := strings.Split(input, ",")

		filterBlocks := regexp.MustCompile(`(\d+)`)
		for _, direction := range directions {
			blocks, _ := strconv.Atoi(filterBlocks.FindAllString(direction, -1)[0])

			if strings.Contains(direction, "R") {
				passenger.Turn("R")
			} else {
				passenger.Turn("L")
			}
			for i := 0; i < blocks; i++ {
				passenger.Walk(1)

				currentPlace := Location{passenger.x, passenger.y}
				if _, keyExists := visited[currentPlace]; !keyExists {
					visited[currentPlace] = true
				} else if !secondSolved {
					fmt.Println("Place with coordinates", currentPlace, "has alreay been visited")
					fmt.Println("Distance from Landing zone:", distanceToHQ(currentPlace.x, currentPlace.y))
					secondSolved = true
				}
			}

		}
	}

	distance := distanceToHQ(passenger.x, passenger.y)
	fmt.Println("Distance from landing zone to HQ: ", distance)
}

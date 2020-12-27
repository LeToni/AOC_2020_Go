package main

import (
	"bufio"
	"os"
)

type position struct {
	posX, posY int
}

const (
	gridSizeX int = 100
	gridSizeY int = 100
)

var (
	grid       = make([][]byte, 100)
	directions = []position{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}
)

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Bytes()
		grid[i] = row
		i++
	}
}

func adjacentNeighborsHaveLightOn(x, y int) int {
	amountNeighbors := 0

	for _, dir := range directions {
		neighbor := neighborAtDirection(x, y, dir)
		if withinGrid(neighbor) && HasLightOn(neighbor) {
			amountNeighbors++
		}
	}

	return amountNeighbors
}

func HasLightOn(p position) bool {
	if withinGrid(p) && grid[p.posX][p.posY] == '#' {
		return true
	} else {
		return false
	}
}

func withinGrid(p position) bool {
	return !(p.posX < 0 || p.posX >= gridSizeX || p.posY < 0 || p.posY >= gridSizeY)
}

func neighborAtDirection(x, y int, p position) position {
	neighbor := position{p.posX + x, p.posY + y}
	return neighbor
}

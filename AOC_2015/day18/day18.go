package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	posX, posY int
}

const (
	gridSizeX int = 100
	gridSizeY int = 100
	maxSteps  int = 100
)

var (
	grid       [gridSizeX][gridSizeY]byte
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
		for j := range row {
			if isCornerLight(i, j) {
				grid[i][j] = '#'
			} else {
				grid[i][j] = row[j]
			}

		}
		i++
	}

	for steps := 0; steps < maxSteps; steps++ {
		var currentGrid [gridSizeX][gridSizeY]byte
		for i := range grid {
			for j := range grid[i] {
				if isCornerLight(i, j) {
					currentGrid[i][j] = '#'
					continue
				}
				neighborsLightOn := adjacentNeighborsHaveLightOn(i, j)
				if grid[i][j] == '#' {
					if neighborsLightOn == 2 || neighborsLightOn == 3 {
						currentGrid[i][j] = '#'
					} else {
						currentGrid[i][j] = '.'
					}
				} else if grid[i][j] == '.' {
					if neighborsLightOn == 3 {
						currentGrid[i][j] = '#'
					} else {
						currentGrid[i][j] = '.'
					}
				}

			}
		}
		grid = currentGrid
	}

	lightsOn := countAmountOfLightsOn()
	fmt.Println("Total lights on: ", lightsOn)

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

func countAmountOfLightsOn() int {
	count := 0

	for _, row := range grid {
		for _, light := range row {
			if light == '#' {
				count = count + 1
			}
		}
	}
	return count
}

func isCornerLight(x, y int) bool {
	if (x == gridSizeX-1 || x == 0) && x == y {
		return true
	} else if x == 0 && y == gridSizeY-1 {
		return true
	} else if x == gridSizeX-1 && y == 0 {
		return true
	} else {
		return false
	}
}

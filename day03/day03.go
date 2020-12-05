package main

import (
	"bufio"
	"fmt"
	"os"
)

type slope struct {
	x int
	y int
}

// Day3 Advent of Code Challenge
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var grid []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	path := slope{x: 3, y: 1}
	answer := traverseMap(grid, path)
	fmt.Println("Answer to first puzzle: ", answer)

	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	factor := 1
	for _, path := range slopes {
		factor = factor * traverseMap(grid, path)
	}

	fmt.Println("Answer to second puzzle: ", factor)
}

func traverseMap(grid []string, path slope) int {

	posX := path.x
	posY := path.y
	countTrees := 0
	for posY < len(grid) {

		if posX > len(grid[posY])-1 {
			posX = posX - len(grid[posY])
		}
		if grid[posY][posX] == '#' {
			countTrees++
		}

		posY = posY + path.y
		posX = posX + path.x
	}
	return countTrees
}

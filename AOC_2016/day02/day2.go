package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

var (
	KeyPad = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	position   = Position{x: 1, y: 1}
	directions string
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		directions = scanner.Text()

		for _, direction := range directions {
			move(direction)
			// fmt.Println(KeyPad[position.x][position.y])
		}
		fmt.Print(KeyPad[position.x][position.y])
	}
	fmt.Println("")
}

func move(direction rune) {
	if direction == 'U' && !(position.x == 0) {
		position.x = position.x - 1
	} else if direction == 'L' && !(position.y == 0) {
		position.y = position.y - 1
	} else if direction == 'D' && !(position.x == 2) {
		position.x = position.x + 1
	} else if direction == 'R' && !(position.y == 2) {
		position.y = position.y + 1
	}
}

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

	AdvancedKeyPad = [5][5]byte{
		{'-', '-', '1', '-', '-'},
		{'-', '2', '3', '4', '-'},
		{'5', '6', '7', '8', '9'},
		{'-', 'A', 'B', 'C', '-'},
		{'-', '-', '9', '-', '-'},
	}

	// position   = Position{x: 1, y: 1}
	position   = Position{x: 0, y: 3}
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
			// moveOnKeyPad(direction)
			moveOnAdvancedKeyPad(direction)
			// fmt.Println(KeyPad[position.x][position.y])
		}
		// fmt.Print(KeyPad[position.x][position.y])
		fmt.Print(string(AdvancedKeyPad[position.x][position.y]))
	}
	fmt.Println("")
}

func moveOnKeyPad(direction rune) {
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

func moveOnAdvancedKeyPad(direction rune) {
	if direction == 'U' && !(position.x == 0) && AdvancedKeyPad[position.x-1][position.y] != '-' {
		position.x = position.x - 1
	} else if direction == 'L' && !(position.y == 0) && AdvancedKeyPad[position.x][position.y-1] != '-' {
		position.y = position.y - 1
	} else if direction == 'D' && !(position.x == len(AdvancedKeyPad)-1) && AdvancedKeyPad[position.x+1][position.y] != '-' {
		position.x = position.x + 1
	} else if direction == 'R' && !(position.y == len(AdvancedKeyPad)-1) && AdvancedKeyPad[position.x][position.y+1] != '-' {
		position.y = position.y + 1
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Triangle struct {
	sideA, sideB, sideC int
}

func (triangle *Triangle) IsTriangle() bool {
	if triangle.sideA+triangle.sideB > triangle.sideC &&
		triangle.sideA+triangle.sideC > triangle.sideB &&
		triangle.sideB+triangle.sideC > triangle.sideA {
		return true
	} else {
		return false
	}
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		input := scanner.Text()

		var (
			a, b, c int
		)
		if n, _ := fmt.Sscanf(input, "  %d  %d  %d", &a, &b, &c); n == 3 {
			triangle := Triangle{a, b, c}
			if triangle.IsTriangle() {
				count++
			}
		} else {
			panic(n)
		}
	}
	fmt.Println("Number of triangles found: ", count)
}

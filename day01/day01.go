package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Day1 Advent of Code Challenge
// Solution
func main() {
	// Open the file
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numbers := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)

	}

	solveDay1FirstPuzzle(numbers)
	solveDay1SecondPuzzle(numbers)
}

func solveDay1FirstPuzzle(numbers []int) {
	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				fmt.Println("Answer to first puzzle: ", i*j)
				return
			}
		}
	}
}

func solveDay1SecondPuzzle(numbers []int) {
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					fmt.Println("Answer to second puzzle: ", i*j*k)
					return
				}
			}
		}
	}
}

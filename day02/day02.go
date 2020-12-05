package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	min      int
	max      int
	letter   string
	password string
}

// Day2 Advent of Code Challenge
// Solution
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	passwords := []*password{}

	regexFilter := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		min, _ := strconv.Atoi(regexFilter.ReplaceAllString(line, "$1"))
		max, _ := strconv.Atoi(regexFilter.ReplaceAllString(line, "$2"))

		passwords = append(passwords, &password{
			min:      min,
			max:      max,
			letter:   regexFilter.ReplaceAllString(line, "$3"),
			password: regexFilter.ReplaceAllString(line, "$4"),
		})
	}

	solveDay2FirstPuzzle(passwords)
	solveDay2SecondPuzzle(passwords)
}

func solveDay2FirstPuzzle(passwords []*password) {
	counter := 0

	for _, pw := range passwords {
		countLetter := strings.Count(pw.password, pw.letter)
		if pw.min <= countLetter && countLetter <= pw.max {
			counter++
		}
	}

	fmt.Println("Day 2 Solution to Puzzle 1: ", counter)
}

func solveDay2SecondPuzzle(passwords []*password) {
	counter := 0
	for _, pw := range passwords {
		if (string(pw.password[pw.min-1]) == pw.letter) != (string(pw.password[pw.max-1]) == pw.letter) {
			counter++
		}
	}

	fmt.Println("Day 2 Solution to Puzzle 2: ", counter)
}

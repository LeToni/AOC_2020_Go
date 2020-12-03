package main

import (
	"fmt"
	"os"

	days "adventOfCode/days"
)

func main() {
	selectDay(os.Args[1])
}

func selectDay(day string) {
	switch day {
	case "1":
		fmt.Println("Running puzzles from day 1")
		days.Day1()
	case "2":
		fmt.Println("Running puzzle from day 2")
		days.Day2()
	case "3":
		fmt.Println("Running puzzle from day 3")
		days.Day3()
	default:
		fmt.Println("Puzzle for Day ", day, " not available yet")
	}

	fmt.Println("")
}

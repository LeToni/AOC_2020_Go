package main

import (
	"bufio"
	"fmt"
	"os"
)

type rule struct {
	leftSide, rightSide string
}

var (
	rules     []*rule
	molecules = make(map[string]bool)
	molecule  string
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			continue
		}

		var inputLeft, inputRight string
		if n, _ := fmt.Sscanf(input, "%s => %s", &inputLeft, &inputRight); n == 2 {
			inputRule := rule{leftSide: inputLeft, rightSide: inputRight}
			rules = append(rules, &inputRule)
		} else {
			molecule = input
		}
	}
}

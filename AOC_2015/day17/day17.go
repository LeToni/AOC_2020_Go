package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	capacity = 150
)

var (
	containers = []int{}
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
		container, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		containers = append(containers, container)
	}
	sort.Ints(containers)
	total := amountPossibleCombinations(containers, capacity)
	fmt.Println("Amount of possible combos:", total)

	total = amountUniqueCombinations(containers, capacity)
	fmt.Println("Amount of possible uniqe combos:", total)
}

func amountPossibleCombinations(containers []int, target int) int {
	total := 0

	if len(containers) == 2 {
		if containers[0]+containers[1] == target {

			return 1
		} else {
			return 0
		}
	}

	for i := 1; i < len(containers); i++ {
		if containers[0]+containers[i] == target {
			total = total + 1
		}
	}

	return total + amountPossibleCombinations(containers[1:], target) + amountPossibleCombinations(containers[1:], target-containers[0])
}

func amountUniqueCombinations(containers []int, target int) int {
	total := 0

	if len(containers) == 2 {
		if containers[0]+containers[1] == target {

			return 1
		} else {
			return 0
		}
	}

	counter := 1
	for i := 1; i < len(containers); i++ {
		if containers[0] == containers[i] {
			counter = counter + 1
		}
	}

	for i := counter; i < len(containers); i++ {
		if containers[0]+containers[i] == target && containers[i] != containers[i-1] {
			total = total + 1
		}
	}

	return total + amountUniqueCombinations(containers[counter:], target) + amountUniqueCombinations(containers[1:], target-containers[0])
}

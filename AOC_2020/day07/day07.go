package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Bag struct {
	color    string
	contents map[string]int
}

func NewBag(bagColor string) *Bag {
	return &Bag{
		color:    bagColor,
		contents: make(map[string]int),
	}
}

func (bag *Bag) ContainsBag(color string) bool {

	eval := false
	if len(bag.contents) == 0 {
		return false
	}
	if _, ok := bag.contents[color]; ok {
		return true
	}

	for k := range bag.contents {
		eval = eval || bags[k].ContainsBag(color)
	}

	return eval
}

var (
	bags = make(map[string]*Bag)
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	regFilterBag := regexp.MustCompile(`^(\w+\s\w+) bags contain`)
	regFilterContent := regexp.MustCompile(`(?P<amount>\d+)\s(?P<bag>\w+\s\w+)\sbags?.|,`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		bagColor := regFilterBag.FindStringSubmatch(input)[1]
		newBag := NewBag(bagColor)

		if n, _ := fmt.Sscan(input, "%s %s bags contain no other bags."); n == 2 {
			continue
		}

		for _, content := range regFilterContent.FindAllString(input, -1) {
			bag := regFilterContent.ReplaceAllString(content, "$bag")
			amount, _ := strconv.Atoi(regFilterContent.ReplaceAllString(content, "$amount"))
			newBag.contents[bag] = amount
		}

		bags[bagColor] = newBag
	}

	TaskOne()
	TaskTwo("shiny gold")
}

func TaskOne() {
	count := 0

	for _, bag := range bags {
		if bag.ContainsBag("shiny gold") {
			count++
		}
	}

	fmt.Println("Number of bags that contain eventually a shiny gold bag:", count)
}

func TaskTwo(bag string) {
	amount := AmountOfBagsIn(bags[bag]) - 1

	fmt.Println("Amount of bags that fit in", bag, "bag:", amount)
}

func AmountOfBagsIn(bag *Bag) int {
	amount := 1

	if len(bag.contents) == 0 {
		return 1
	}
	for k := range bag.contents {
		// eval = eval || bags[k].ContainsBag(color)
		amount = amount + bag.contents[k]*AmountOfBagsIn(bags[k])
	}
	return amount
}

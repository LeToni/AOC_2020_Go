package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Intervall struct {
	lowerBound, upperBound int
}

func (inter *Intervall) withinBoundary(number int) bool {
	if inter.lowerBound <= number && number <= inter.upperBound {
		return true
	} else {
		return false
	}
}

type TicketField struct {
	name  string
	rule1 Intervall
	rule2 Intervall
	index int
	pos map[int]bool
	pos   map[int]bool
}

func (tf *TicketField) isValid(number int) bool {

	if tf.rule1.withinBoundary(number) || tf.rule2.withinBoundary(number) {
		return true
	} else {
		return false
	}
}

type Ticket struct {
	numbers []int
}

func (ticket *Ticket) CalculateError() {
	for _, number := range ticket.numbers {
		validNumber := false
		for _, tf := range ticketFields {
		for index, tf := range ticketFields {
			if tf.isValid(number) {
				validNumber = true
			} else {
				tf.pos[index] = false
			}
		}
		if !validNumber {
			errorRate = errorRate + number
		}
	}

}

var (
	ticketFields  = []TicketField{}
	ownTicket     Ticket
	nearByTickets = []Ticket{}
	errorRate     int
)

func main() {
	readInput()
	for _, nbt := range nearByTickets {
		nbt.CalculateError()
	}
	fmt.Println("Task 1 -> Scanning error rate for nearby tickets:", errorRate)
}

func readInput() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(file), "\n\n")

	regFieldName := regexp.MustCompile(`(\w+)(\s\w+)*`)
	regBoundaries := regexp.MustCompile(`(\d+)-(\d+)\sor\s(\d+)-(\d+)`)
	for _, ticket := range strings.Split(string(content[0]), "\n") {
		fieldName := regFieldName.FindString(ticket)
		boundaries := regBoundaries.FindAllStringSubmatch(ticket, -1)[0]

		min1Intervall, _ := strconv.Atoi(boundaries[1])
		max1Intervall, _ := strconv.Atoi(boundaries[2])
		min2Intervall, _ := strconv.Atoi(boundaries[3])
		max2Intervall, _ := strconv.Atoi(boundaries[4])
		rule1 := Intervall{min1Intervall, max1Intervall}
		rule2 := Intervall{min2Intervall, max2Intervall}

		newTicketField := TicketField{name: fieldName, rule1: rule1, rule2: rule2}
		ticketFields = append(ticketFields, newTicketField)
	}

	// reading own ticket
	for i, ticket := range strings.Split(string(content[1]), "\n") {
		if i == 0 {
			continue
		}
		numbers := strings.Split(ticket, ",")

		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			ownTicket.numbers = append(ownTicket.numbers, n)
		}
	}

	// Reading nearby tickets
	for i, ticket := range strings.Split(string(content[2]), "\n") {
		if i == 0 {
			continue
		}
		var nearbyTicket Ticket
		numbers := strings.Split(ticket, ",")

		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			nearbyTicket.numbers = append(nearbyTicket.numbers, n)
		}

		nearByTickets = append(nearByTickets, nearbyTicket)
	}
}

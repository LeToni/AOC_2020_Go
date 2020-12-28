package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type rule struct {
	leftSide, rightSide string
}
type electron struct {
	rightSide string
}

var (
	rules     []*rule
	molecules = make(map[string]bool)
	molecule  string
	electrons []*electron
	element   = make(map[string]bool)
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
		if n, _ := fmt.Sscanf(input, "e => %s", &inputRight); n == 1 {
			e := electron{rightSide: inputRight}
			electrons = append(electrons, &e)
		} else if n, _ := fmt.Sscanf(input, "%s => %s", &inputLeft, &inputRight); n == 2 {
			inputRule := rule{leftSide: inputLeft, rightSide: inputRight}
			rules = append(rules, &inputRule)
			element[inputLeft] = true
		} else {
			molecule = input
		}
	}
	CreateMolecules()
	fmt.Println("Amount of distinct molecules that can be created:", len(molecules))
	steps := RecreateMolecule()
	fmt.Println("Steps required to produce molecule:", steps)
}

func CreateMolecules() {
	for _, r := range rules {
		regex := regexp.MustCompile(r.leftSide)

		if !regex.MatchString(molecule) {
			continue
		}
		particles := regex.FindAllStringIndex(molecule, -1)

		for _, particle := range particles {
			newMolecule := make([]byte, len(molecule))
			copy(newMolecule, molecule)
			newMolecule = append(newMolecule[:particle[0]], append([]byte(r.rightSide), newMolecule[particle[1]:]...)...)
			molecules[string(newMolecule)] = true
		}
	}
}

func RecreateMolecule() int {
	steps := 0
	e := molecule
	regFilterTerminals := regexp.MustCompile("Rn|Ar")
	steps = steps + len(regFilterTerminals.FindAllStringIndex(e, -1))
	e = regFilterTerminals.ReplaceAllString(e, "")

	regFilterTerminals = regexp.MustCompile("Y")
	steps = steps + len(regFilterTerminals.FindAllStringIndex(e, -1))*2
	e = regFilterTerminals.ReplaceAllString(e, "")

	for _, r := range rules {
		filterReg := regexp.MustCompile(r.rightSide)

		if !filterReg.MatchString(e) {
			continue
		}

		steps = steps + len(filterReg.FindAllStringIndex(e, -1))
		e = strings.ReplaceAll(e, r.rightSide, "")
	}

	for _, r := range rules {
		filterReg := regexp.MustCompile(r.leftSide)

		if !filterReg.MatchString(e) {
			continue
		}

		steps = steps + len(filterReg.FindAllStringIndex(e, -1))
		e = strings.ReplaceAll(e, r.leftSide, "")
	}
	return steps - len(element) - 1
}

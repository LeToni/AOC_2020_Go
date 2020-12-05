package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Day4 Advent of Code Challenge
func main() {
	// Open the file
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	passports := []map[string]string{}
	passport := map[string]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			passports = append(passports, passport)
			passport = map[string]string{}
		} else {
			for _, subString := range strings.Split(text, " ") {
				stringPair := strings.Split(subString, ":")
				passport[stringPair[0]] = stringPair[1]
			}
		}
	}
	passports = append(passports, passport)

	answer := eval(passports)
	fmt.Println("Answer to first puzzle: ", answer)

	answer = evalStrictly(passports)
	fmt.Println("Answer to first puzzle: ", answer)
}

func eval(passports []map[string]string) int {
	count := 0
	for _, passport := range passports {
		if isValid(passport) {
			count++
		}
	}
	return count
}

func isValid(passport map[string]string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range fields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func evalStrictly(passports []map[string]string) int {
	count := 0
	for _, passport := range passports {
		if isValid(passport) {
			if isFormatValid(passport) {
				count++
			}
		}
	}
	return count
}

func isFormatValid(passport map[string]string) bool {

	return byrValid(passport["byr"]) && iyrValid(passport["iyr"]) && eyrValid(passport["eyr"]) && hgtValid(passport["hgt"]) && hclValid(passport["hcl"]) && byrValid(passport["byr"]) && eclValid(passport["ecl"]) && pidValid(passport["pid"])
}

func byrValid(byr string) bool {
	byrInt, err := strconv.Atoi(byr)
	if err != nil || byrInt < 1920 || byrInt > 2002 {
		return false
	}
	return true
}

func iyrValid(iyr string) bool {
	iyrInt, err := strconv.Atoi(iyr)
	if err != nil || iyrInt < 2010 || iyrInt > 2020 {
		return false
	}
	return true
}

func eyrValid(eyr string) bool {
	eyrInt, err := strconv.Atoi(eyr)
	if err != nil || eyrInt < 2020 || eyrInt > 2030 {
		return false
	}
	return true
}

func hgtValid(hgt string) bool {
	if strings.HasSuffix(hgt, "cm") {
		if no, err := strconv.Atoi(strings.TrimSuffix(hgt, "cm")); err != nil || no < 150 || no > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		if no, err := strconv.Atoi(strings.TrimSuffix(hgt, "in")); err != nil || no < 59 || no > 76 {
			return false
		}
	} else {
		return false
	}
	return true
}

func hclValid(hcl string) bool {
	return regexp.MustCompile(`^#[a-f0-9]{6}$`).Match([]byte(hcl))
}
func pidValid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	if id, err := strconv.Atoi(pid); err != nil || id == 0 {
		return false
	}
	return true
}

func eclValid(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
}

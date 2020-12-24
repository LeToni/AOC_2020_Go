package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func extractNumbers(jsonData interface{}) []int {
	numbers := []int{}

	switch data := jsonData.(type) {
	case []interface{}:
		for _, value := range data {
			numbers = append(numbers, extractNumbers(value)...)
		}
	case map[string]interface{}:
		containsRed := false

		for _, key := range data {
			if str, ok := key.(string); ok && str == "red" {
				containsRed = true
				break
			}
		}
		if !containsRed {
			for _, value := range data {
				numbers = append(numbers, extractNumbers(value)...)
			}
		}

	case float64:
		numbers = append(numbers, int(data))
	}

	return numbers
}

func main() {
	document, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	json.Unmarshal(document, &data)

	foundNumbers := extractNumbers(data)

	total := 0
	for _, number := range foundNumbers {
		total = total + number
	}

	fmt.Println("Result ", total)
}

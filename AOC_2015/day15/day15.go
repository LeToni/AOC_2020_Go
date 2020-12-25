package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

const (
	teaspoons   int = 100
	maxCalories int = 500
)

var (
	ingredients = []Ingredient{}
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ingredient := scanner.Text()

		ingredient = strings.ReplaceAll(ingredient, ",", "")
		ingredient = strings.ReplaceAll(ingredient, ":", "")
		var (
			name                                            string
			capacity, durability, flavor, texture, calories int
		)
		if n, _ := fmt.Sscanf(ingredient, "%s capacity %d durability %d flavor %d texture %d calories %d", &name, &capacity, &durability, &flavor, &texture, &calories); n == 6 {
			newIngredient := Ingredient{name, capacity, durability, flavor, texture, calories}
			ingredients = append(ingredients, newIngredient)
		} else {
			panic(ingredient)
		}
	}

	result := calculateBestRatio()
	fmt.Printf("Result with the best score: %d\n", result)
	result = calculateBestCaloriesRatio()
	fmt.Printf("Result with the best score and max calories: %d\n", result)
}

func calculateBestCaloriesRatio() int {
	bestScore := 0

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100-a; b++ {
			for c := 0; c <= 100-a-b; c++ {
				d := 100 - (a + b + c)
				score, calories := calculateScore(a, b, c, d)
				if score > bestScore && calories <= maxCalories {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}

func calculateBestRatio() int {
	bestScore := 0

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100-a; b++ {
			for c := 0; c <= 100-a-b; c++ {
				d := 100 - (a + b + c)
				score, _ := calculateScore(a, b, c, d)
				if score > bestScore {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}

func calculateScore(scores ...int) (int, int) {

	var (
		capacity, durability, flavor, texture, calories int
	)

	for i := 0; i < len(scores); i++ {
		capacity = capacity + ingredients[i].capacity*scores[i]
		durability = durability + ingredients[i].durability*scores[i]
		flavor = flavor + ingredients[i].flavor*scores[i]
		texture = texture + ingredients[i].texture*scores[i]
		calories = calories + ingredients[i].calories*scores[i]
	}

	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0, 0
	}

	score := capacity * durability * flavor * texture

	return score, calories
}

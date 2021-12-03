package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type product struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func read() []product {
	lines := aoc.ReadAllInput()
	var products []product

	for _, line := range lines {
		if line == "" {
			continue
		}
		sp := strings.Split(line, " (contains ")
		i := strings.Split(sp[0], " ")
		ingredients := make(map[string]bool, len(i))
		for _, v := range i {
			ingredients[v] = true
		}
		a := strings.Split(sp[1][:len(sp[1])-1], ", ")
		allergens := make(map[string]bool, len(a))
		for _, v := range a {
			allergens[v] = true
		}
		products = append(products, product{ingredients: ingredients, allergens: allergens})
	}

	return products
}

func process(data []product) int {
	possibleAllergens := make(map[string]map[string]bool) // allergen -> possible ingredients

	// collect all possible allergens
	for _, product := range data {
		for a := range product.allergens {
			ingredients := possibleAllergens[a]
			if ingredients == nil {
				ingredients = make(map[string]bool)
				possibleAllergens[a] = ingredients
			}
			for i := range product.ingredients {
				ingredients[i] = true
			}
		}
	}

	// remove ingredients
	for _, p := range data {
		for allergen := range p.allergens {
			possibleIngredients := possibleAllergens[allergen]
			for i := range possibleIngredients {
				if !p.ingredients[i] {
					delete(possibleIngredients, i)
				}
			}
		}
	}

	ingredientsWithAllergens := make(map[string]bool)
	for _, ingredients := range possibleAllergens {
		for i := range ingredients {
			ingredientsWithAllergens[i] = true
		}
	}

	// count ingredients
	count := 0
	for _, p := range data {
		for i := range p.ingredients {
			if !ingredientsWithAllergens[i] {
				count++
			}
		}
	}

	return count
}

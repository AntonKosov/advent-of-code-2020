package main

import (
	"fmt"
	"sort"
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

func process(data []product) string {
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

	// find ingredients and allergens
	foundIngredients := make(map[string]string) // ingredient -> allergen
	foundAllergens := make(map[string]string)   // allergen -> ingredient
	var deleteIngredient func(allergen string, ingredient string)
	deleteIngredient = func(allergen string, ingredient string) {
		ingredients := possibleAllergens[allergen]
		if !ingredients[ingredient] {
			return
		}
		delete(ingredients, ingredient)
		if len(ingredients) != 1 {
			return
		}
		// the single value is the ingredient
		for i := range ingredients {
			foundAllergens[allergen] = i
			foundIngredients[i] = allergen
			for pa := range possibleAllergens {
				deleteIngredient(pa, i)
			}
			break
		}
	}
	for _, p := range data {
		for allergen := range p.allergens {
			possibleIngredients := possibleAllergens[allergen]
			for i := range possibleIngredients {
				if !p.ingredients[i] {
					deleteIngredient(allergen, i)
				}
			}
		}
	}

	// sort allergens
	var allergens []string
	for a := range foundAllergens {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)

	// build the list of ingredients
	var result strings.Builder
	for _, a := range allergens {
		if result.Len() > 0 {
			result.WriteRune(',')
		}
		result.WriteString(foundAllergens[a])
	}

	return result.String()
}

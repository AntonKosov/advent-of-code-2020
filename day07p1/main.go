package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type included map[string]bool

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (map[string]included, error) {
	result := make(map[string]included)
	handler := func(line string) error {
		splitted := strings.Split(line, " bags contain ")
		bigBag := splitted[0]
		smallerBags := strings.Split(splitted[1], ", ")
		for _, smallerBag := range smallerBags {
			parts := strings.Split(smallerBag, " ")
			color := fmt.Sprintf("%v %v", parts[1], parts[2])
			biggerBags, ok := result[color]
			if !ok {
				biggerBags = included{}
				result[color] = biggerBags
			}
			biggerBags[bigBag] = true
		}

		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

func process(data map[string]included) int {
	contain := make(map[string]bool)
	check := []string{"shiny gold"}
	for len(check) > 0 {
		color := check[0]
		if len(check) > 1 {
			check = check[1:]
		} else {
			check = check[:0]
		}
		for biggerBag := range data[color] {
			if !contain[biggerBag] {
				contain[biggerBag] = true
				check = append(check, biggerBag)
			}
		}
	}

	return len(contain)
}

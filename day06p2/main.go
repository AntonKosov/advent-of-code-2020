package main

import (
	"fmt"
	"log"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type group []string

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([]group, error) {
	result := []group{}
	var g group
	handler := func(line string) error {
		if line == "" {
			result = append(result, g)
			g = nil
			return nil
		}
		if g == nil {
			g = group{}
		}
		g = append(g, line)
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	result = append(result, g)
	return result, nil
}

func process(data []group) int {
	sum := 0
	for _, group := range data {
		m := make(map[rune]int)
		for _, questions := range group {
			for _, question := range questions {
				m[question]++
			}
		}
		for _, c := range m {
			if c == len(group) {
				sum++
			}
		}
	}

	return sum
}

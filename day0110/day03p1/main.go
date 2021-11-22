package main

import (
	"fmt"
	"log"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type square rune

const (
	free square = 'f'
	tree square = 't'
)

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func process(data [][]square) int {
	count := 0
	col := 0
	for i := 1; i < len(data); i++ {
		row := data[i]
		col = (col + 3) % len(row)
		if row[col] == tree {
			count++
		}
	}
	return count
}

func read() ([][]square, error) {
	result := [][]square{}
	handler := func(line string) error {
		row := make([]square, len(line))
		for i, s := range line {
			v := tree
			if s == '.' {
				v = free
			}
			row[i] = v
		}
		result = append(result, row)
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

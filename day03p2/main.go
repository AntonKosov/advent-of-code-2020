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
	slopes := []struct {
		dc int
		dr int
	}{
		{dc: 1, dr: 1},
		{dc: 3, dr: 1},
		{dc: 5, dr: 1},
		{dc: 7, dr: 1},
		{dc: 1, dr: 2},
	}

	mul := 1
	for _, s := range slopes {
		mul *= countTrees(data, s.dc, s.dr)
	}

	return mul
}

func countTrees(data [][]square, dc, dr int) int {
	width := len(data[0])
	count := 0
	col := 0
	row := 0
	for {
		row += dr
		if row >= len(data) {
			return count
		}
		col = (col + dc) % width
		line := data[row]
		if line[col] == tree {
			count++
		}
	}
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

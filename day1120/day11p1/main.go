package main

import (
	"fmt"
	"log"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type cell rune

const (
	floor    cell = '.'
	free     cell = 'L'
	occupied cell = '#'
)

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([][]cell, error) {
	result := [][]cell{}
	handler := func(line string) error {
		row := make([]cell, len(line))
		for i, c := range line {
			row[i] = cell(c)
		}
		result = append(result, row)

		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

type change struct {
	row int
	col int
	c   cell
}

func process(data [][]cell) int {
	changes := make(chan change, len(data)*len(data[0]))
	iterations := 0
	for {
		for r, row := range data {
			for c, cell := range row {
				if cell == floor {
					continue
				}
				adj := countAdj(data, r, c)
				if adj == 0 && cell != occupied {
					changes <- change{row: r, col: c, c: occupied}
				} else if adj >= 4 && cell != free {
					changes <- change{row: r, col: c, c: free}
				}
			}
		}
		if len(changes) == 0 {
			return occupiedSeats(data)
		}
		iterations++
		for len(changes) > 0 {
			c := <-changes
			data[c.row][c.col] = c.c
		}
	}
}

func occupiedSeats(data [][]cell) int {
	sum := 0
	for _, row := range data {
		for _, c := range row {
			if c == occupied {
				sum++
			}
		}
	}
	return sum
}

func countAdj(data [][]cell, row, column int) int {
	sum := 0
	for r := aoc.Max(0, row-1); r < aoc.Min(len(data), row+2); r++ {
		hLine := data[r]
		for c := aoc.Max(0, column-1); c < aoc.Min(len(hLine), column+2); c++ {
			if r == row && c == column {
				continue
			}
			if hLine[c] == occupied {
				sum++
			}
		}
	}
	return sum
}

package main

import (
	"fmt"
	"log"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

const maxOccupiedSeats = 5

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
				} else if adj >= maxOccupiedSeats && cell != free {
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
	fo := func(dr, dc int) int {
		r, c := row, column
		for {
			r += dr
			c += dc
			if r < 0 || r >= len(data) || c < 0 || c >= len(data[0]) {
				return 0
			}
			switch data[r][c] {
			case occupied:
				return 1
			case free:
				return 0
			}
		}
	}

	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			sum += fo(i, j)
			if sum >= maxOccupiedSeats {
				return sum
			}
		}
	}
	return sum
}

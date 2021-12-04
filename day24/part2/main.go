package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

const (
	stepE  = "e"
	stepSE = "se"
	stepSW = "sw"
	stepW  = "w"
	stepNW = "nw"
	stepNE = "ne"
)

type path []string

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() []path {
	lines := aoc.ReadAllInput()
	var result []path

	allSteps := []string{stepNE, stepNW, stepSE, stepSW, stepE, stepW}
	for _, line := range lines {
		if line == "" {
			break
		}
		ss := make(path, len(line))
		for line != "" {
			for _, s := range allSteps {
				if strings.HasPrefix(line, s) {
					ss = append(ss, s)
					line = line[len(s):]
					break
				}
			}
		}

		result = append(result, ss)
	}

	return result
}

func process(paths []path) int {
	adjacent := map[string]aoc.Vector2{
		stepE:  aoc.NewVector2(2, 0),
		stepW:  aoc.NewVector2(-2, 0),
		stepNE: aoc.NewVector2(1, -1),
		stepNW: aoc.NewVector2(-1, -1),
		stepSE: aoc.NewVector2(1, 1),
		stepSW: aoc.NewVector2(-1, 1),
	}

	grid := make(map[aoc.Vector2]bool) // coordinate -> isBlack
	for _, path := range paths {
		targetTile := aoc.NewVector2(0, 0)
		for _, s := range path {
			targetTile = targetTile.Add(adjacent[s])
		}

		if grid[targetTile] {
			delete(grid, targetTile)
		} else {
			grid[targetTile] = true
		}
	}

	countBlackAround := func(pos aoc.Vector2) int {
		sum := 0
		for _, a := range adjacent {
			if grid[pos.Add(a)] {
				sum++
			}
		}
		return sum
	}

	for i := 0; i < 100; i++ {
		nextGrid := make(map[aoc.Vector2]bool, len(grid)*2)
		for blackPos := range grid {
			blackAround := countBlackAround(blackPos)
			if blackAround == 1 || blackAround == 2 {
				nextGrid[blackPos] = true
			}
			for _, adjOffset := range adjacent {
				adjPos := blackPos.Add(adjOffset)
				if !grid[adjPos] && countBlackAround(adjPos) == 2 {
					nextGrid[adjPos] = true
				}
			}
		}
		grid = nextGrid
	}

	return len(grid)
}

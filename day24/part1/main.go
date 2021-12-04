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
	stepToDir := map[string]aoc.Vector2{
		stepE:  aoc.NewVector2(2, 0),
		stepW:  aoc.NewVector2(-2, 0),
		stepNE: aoc.NewVector2(1, -1),
		stepNW: aoc.NewVector2(-1, -1),
		stepSE: aoc.NewVector2(1, 1),
		stepSW: aoc.NewVector2(-1, 1),
	}

	m := make(map[aoc.Vector2]bool) // coordinate -> isBlack
	for _, path := range paths {
		targetTile := aoc.NewVector2(0, 0)
		for _, s := range path {
			targetTile = targetTile.Add(stepToDir[s])
		}

		m[targetTile] = !m[targetTile]
	}

	black := 0
	for _, c := range m {
		if c {
			black++
		}
	}

	return black
}

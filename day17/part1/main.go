package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() map[aoc.Vector3]struct{} {
	lines := aoc.ReadAllInput()

	data := map[aoc.Vector3]struct{}{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				data[aoc.NewVector3(x, y, 0)] = struct{}{}
			}
		}
	}

	return data
}

func process(data map[aoc.Vector3]struct{}) int {
	for i := 0; i < 6; i++ {
		temp := make(map[aoc.Vector3]struct{}, len(data))
		free := make(map[aoc.Vector3]int, len(data))
		for v := range data {
			n := countNeighbors(data, v, free)
			if n == 2 || n == 3 {
				temp[v] = struct{}{}
			}
		}
		for v, n := range free {
			if n == 3 {
				temp[v] = struct{}{}
			}
		}
		data = temp
	}

	return len(data)
}

func countNeighbors(data map[aoc.Vector3]struct{}, v aoc.Vector3, free map[aoc.Vector3]int) int {
	sum := 0
	for x := v.X - 1; x <= v.X+1; x++ {
		for y := v.Y - 1; y <= v.Y+1; y++ {
			for z := v.Z - 1; z <= v.Z+1; z++ {
				n := aoc.NewVector3(x, y, z)
				if v == n {
					continue
				}
				if _, ok := data[n]; ok {
					sum++
				} else {
					free[n]++
				}
			}
		}
	}
	return sum
}

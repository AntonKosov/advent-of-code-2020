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

func read() map[aoc.Vector4]struct{} {
	lines := aoc.ReadAllInput()

	data := map[aoc.Vector4]struct{}{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				data[aoc.NewVector4(x, y, 0, 0)] = struct{}{}
			}
		}
	}

	return data
}

func process(data map[aoc.Vector4]struct{}) int {
	for i := 0; i < 6; i++ {
		temp := make(map[aoc.Vector4]struct{}, len(data))
		free := make(map[aoc.Vector4]int, len(data))
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

func countNeighbors(data map[aoc.Vector4]struct{}, v aoc.Vector4, free map[aoc.Vector4]int) int {
	sum := 0
	for x := v.X - 1; x <= v.X+1; x++ {
		for y := v.Y - 1; y <= v.Y+1; y++ {
			for z := v.Z - 1; z <= v.Z+1; z++ {
				for w := v.W - 1; w <= v.W+1; w++ {
					n := aoc.NewVector4(x, y, z, w)
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
	}
	return sum
}

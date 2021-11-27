package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() [][]string {
	lines := aoc.ReadAllInput()

	var data [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		data = append(data, strings.Split(line, " "))
	}

	return data
}

func process(data [][]string) int {
	sum := 0

	for _, ex := range data {
		v, _ := calcEx(ex, 0, len(ex)-1)
		sum += v
	}

	return sum
}

func calcEx(ex []string, start, end int) (result int, next int) {
	result, next = getTokenValue(ex, start, end)

	for next <= end {
		operation := ex[next]
		var v int
		v, next = getTokenValue(ex, next+1, end)
		switch operation {
		case "*":
			result *= v
		case "+":
			result += v
		default:
			panic("unknown operation")
		}
	}

	return result, next
}

func getTokenValue(ex []string, start, end int) (result, next int) {
	if ex[start] != "(" {
		return aoc.StrToInt(ex[start]), start + 1
	}

	op := 1
	for i := start + 1; i <= end; i++ {
		switch ex[i] {
		case "(":
			op++
		case ")":
			op--
			if op == 0 {
				v, _ := calcEx(ex, start+1, i-1)
				return v, i + 1
			}
		}
	}

	panic("wrong format")
}

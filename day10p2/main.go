package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([]int, error) {
	result := []int{}
	handler := func(line string) error {
		number, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		result = append(result, number)

		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

func process(data []int) int {
	data = append(data, 0)
	sort.Ints(data)
	cache := make(map[int]int, len(data))
	return countOptons(data, 0, cache)
}

func countOptons(data []int, start int, cache map[int]int) int {
	if start == len(data)-1 {
		return 1
	}
	current := data[start]
	count := 0
	for i := start + 1; i < start+4 && i < len(data); i++ {
		next := data[i]
		if next-current > 3 {
			break
		}
		if c, ok := cache[i]; ok {
			count += c
			continue
		}
		count += countOptons(data, i, cache)
	}

	cache[start] = count

	return count
}

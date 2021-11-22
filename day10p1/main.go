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
	sort.Ints(data)
	one := 0
	three := 1
	prev := 0
	for _, j := range data {
		if j-prev == 1 {
			one++
		} else {
			three++
		}
		prev = j
	}
	return one * three
}

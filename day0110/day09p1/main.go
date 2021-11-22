package main

import (
	"fmt"
	"log"
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
	const preamble = 25
	m := make(map[int]int, len(data))
	for i := 0; i < preamble; i++ {
		m[data[i]] = 1
	}
	for i := preamble; i < len(data); i++ {
		n := data[i]
		found := false
		for k := range m {
			if m[n-k] > 0 {
				found = true
				break
			}
		}
		if !found {
			return n
		}
		m[n]++
		outdatedNumber := data[i-preamble]
		if m[outdatedNumber] > 1 {
			m[outdatedNumber]--
		} else {
			delete(m, outdatedNumber)
		}
	}
	return 0
}

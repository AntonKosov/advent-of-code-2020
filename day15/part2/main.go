package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	var result []int
	handler := func(line string) error {
		values := strings.Split(line, ",")
		for _, s := range values {
			v, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			result = append(result, v)
		}
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

func process(data []int) int {
	spokenTurns := make(map[int]int, len(data))
	prevTurns := make(map[int]int, len(data))
	for i, v := range data {
		prevTurns[v] = i + 1
	}
	lastNumber := data[len(data)-1]
	isNew := true
	const targetTurn = 30_000_000
	for i := len(data) + 1; i <= targetTurn; i++ {
		if isNew {
			lastNumber = 0
		} else {
			lastNumber = spokenTurns[lastNumber] - prevTurns[lastNumber]
		}
		spokenTurn, wasSpoken := spokenTurns[lastNumber]
		_, wasPrevious := prevTurns[lastNumber]
		if wasSpoken {
			prevTurns[lastNumber] = spokenTurn
		}
		isNew = !(wasSpoken || wasPrevious)
		spokenTurns[lastNumber] = i
		if i%(targetTurn/100) == 0 {
			fmt.Printf(">>> %v%%\n", i*100/targetTurn)
		}
	}
	return lastNumber
}

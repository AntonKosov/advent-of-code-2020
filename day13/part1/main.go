package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type schedular struct {
	time  int
	buses []int
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (schedular, error) {
	result := schedular{time: -1}
	handler := func(line string) error {
		if result.time < 0 {
			time, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			result.time = time
			return nil
		}
		buses := strings.Split(line, ",")
		result.buses = make([]int, 0)
		for _, b := range buses {
			if b == "x" {
				continue
			}
			value, err := strconv.Atoi(b)
			if err != nil {
				return err
			}
			result.buses = append(result.buses, value)
		}
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return schedular{}, err
	}
	return result, nil
}

func process(data schedular) int {
	earliestBus := 0
	earliestTime := math.MaxInt
	for _, b := range data.buses {
		nextTime := data.time - (data.time % b) + b
		if nextTime < earliestTime {
			earliestBus = b
			earliestTime = nextTime
		}
	}
	return earliestBus * (earliestTime - data.time)
}

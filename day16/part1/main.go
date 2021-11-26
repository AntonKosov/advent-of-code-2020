package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type tickets struct {
	validValues   map[int]bool
	nearbyTickets []int
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() tickets {
	lines := aoc.ReadAllInput()

	tickets := tickets{validValues: make(map[int]bool)}

	addValidValues := func(tokens ...string) {
		for _, t := range tokens {
			numbers := aoc.StrToInts(t, "-")
			for i := numbers[0]; i <= numbers[1]; i++ {
				tickets.validValues[i] = true
			}
		}
	}

	i := 0
	for ; ; i++ {
		line := lines[i]
		if line == "" {
			break
		}
		s := strings.Split(line, " ")
		addValidValues(s[len(s)-3], s[len(s)-1])
	}

	i++
	for ; ; i++ {
		line := lines[i]
		if line == "" {
			break
		}
	}

	i += 2
	for ; ; i++ {
		line := lines[i]
		if line == "" {
			break
		}
		tickets.nearbyTickets = append(tickets.nearbyTickets, aoc.StrToInts(line, ",")...)
	}

	return tickets
}

func process(data tickets) int {
	sum := 0
	for _, v := range data.nearbyTickets {
		if !data.validValues[v] {
			sum += v
		}
	}
	return sum
}

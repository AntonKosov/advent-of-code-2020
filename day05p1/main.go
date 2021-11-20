package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type seat struct {
	row    string
	column string
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func parseInt(s string, zero, one rune) int {
	s = strings.ReplaceAll(s, string(zero), "0")
	s = strings.ReplaceAll(s, string(one), "1")
	v, err := strconv.ParseInt(s, 2, 8)
	if err != nil {
		panic("cannot parse")
	}
	return int(v)
}

func process(data []seat) int {
	maxId := 0
	for _, s := range data {
		id := parseInt(s.row, 'F', 'B')*8 + parseInt(s.column, 'L', 'R')
		maxId = aoc.Max(maxId, id)
	}

	return maxId
}

func read() ([]seat, error) {
	result := []seat{}
	handler := func(line string) error {
		result = append(result, seat{
			row:    line[:7],
			column: line[7:],
		})
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

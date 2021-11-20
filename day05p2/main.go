package main

import (
	"fmt"
	"log"
	"math"
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
	occupied := make(map[int]bool, len(data))
	minId, maxId := math.MaxInt, 0
	for _, s := range data {
		id := parseInt(s.row, 'F', 'B')*8 + parseInt(s.column, 'L', 'R')
		occupied[id] = true
		minId = aoc.Min(minId, id)
		maxId = aoc.Max(maxId, id)
	}

	for id := minId; id < maxId; id++ {
		if !occupied[id+1] {
			return id + 1
		}
	}

	panic("seat not found")
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

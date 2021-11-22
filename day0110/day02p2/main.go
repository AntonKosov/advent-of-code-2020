package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type password struct {
	pos1 int
	pos2 int
	char byte
	psw  string
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func process(data []password) int {
	count := 0
	for _, p := range data {
		if (p.psw[p.pos1-1] == p.char) != (p.psw[p.pos2-1] == p.char) {
			count++
		}
	}
	return count
}

func read() ([]password, error) {
	result := []password{}
	handler := func(line string) error {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return fmt.Errorf("wrong format")
		}
		count := strings.Split(parts[0], "-")
		if len(count) != 2 {
			return fmt.Errorf("wrong format")
		}
		pos1, err := strconv.Atoi(count[0])
		if err != nil {
			return err
		}
		pos2, err := strconv.Atoi(count[1])
		if err != nil {
			return err
		}
		result = append(result, password{
			pos1: pos1,
			pos2: pos2,
			char: parts[1][0],
			psw:  parts[2],
		})
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

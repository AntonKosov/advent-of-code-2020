package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	min  int
	max  int
	char string
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
		c := strings.Count(p.psw, p.char)
		if c >= p.min && c <= p.max {
			count++
		}
	}
	return count
}

func read() ([]password, error) {
	result := []password{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return nil, fmt.Errorf("wrong format")
		}
		count := strings.Split(parts[0], "-")
		if len(count) != 2 {
			return nil, fmt.Errorf("wrong format")
		}
		min, err := strconv.Atoi(count[0])
		if err != nil {
			return nil, err
		}
		max, err := strconv.Atoi(count[1])
		if err != nil {
			return nil, err
		}
		result = append(result, password{
			min:  min,
			max:  max,
			char: string(parts[1][0]),
			psw:  parts[2],
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

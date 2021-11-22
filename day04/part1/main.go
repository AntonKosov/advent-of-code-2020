package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type passport map[string]string

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func process(data []passport) int {
	valid := 0
	for _, pass := range data {
		vf := len(pass)
		if _, ok := pass["cid"]; ok {
			vf--
		}
		if vf == 7 {
			valid++
		}
	}

	return valid
}

func read() ([]passport, error) {
	result := []passport{}
	var pass map[string]string
	handler := func(line string) error {
		if line == "" {
			result = append(result, pass)
			pass = nil
			return nil
		}
		if pass == nil {
			pass = make(map[string]string, 8)
		}
		fields := strings.Split(line, " ")
		for _, f := range fields {
			kv := strings.Split(f, ":")
			pass[kv[0]] = kv[1]
		}
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	result = append(result, pass)
	return result, nil
}

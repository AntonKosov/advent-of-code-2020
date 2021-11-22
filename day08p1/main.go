package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type operation string

const (
	acc operation = "acc"
	nop operation = "nop"
	jmp operation = "jmp"
)

type instruction struct {
	operation operation
	argument  int
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([]instruction, error) {
	result := []instruction{}
	handler := func(line string) error {
		inst := strings.Split(line, " ")
		argument, err := strconv.Atoi(inst[1])
		if err != nil {
			return err
		}
		result = append(result, instruction{
			operation: operation(inst[0]),
			argument:  argument,
		})

		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

func process(data []instruction) int {
	visited := make(map[int]bool)
	accumulator := 0
	currentLine := 0
	for {
		inst := data[currentLine]
		visited[currentLine] = true
		prevAccValue := accumulator
		switch inst.operation {
		case acc:
			accumulator += inst.argument
			currentLine++
		case nop:
			currentLine++
		case jmp:
			currentLine += inst.argument
		}
		if visited[currentLine] {
			return prevAccValue
		}
	}
}

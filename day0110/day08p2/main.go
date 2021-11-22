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
	visited := make(map[int]bool)
	found, acc := process(data, visited, 0, false, 0)
	fmt.Printf("Answer: %v %v\n", found, acc)
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

func processInstruction(inst instruction, accumulator, line int) (nextLine int, newAcc int) {
	switch inst.operation {
	case acc:
		newAcc = accumulator + inst.argument
		nextLine = line + 1
	case nop:
		newAcc = accumulator
		nextLine = line + 1
	case jmp:
		newAcc = accumulator
		nextLine = line + inst.argument
	}
	return nextLine, newAcc
}

var swappableOperations map[operation]operation

func init() {
	swappableOperations = map[operation]operation{
		nop: jmp,
		jmp: nop,
	}
}

func process(data []instruction, visited map[int]bool, line int, changed bool,
	accumulator int,
) (int, bool) {
	if visited[line] {
		return 0, false
	}
	if line >= len(data) {
		return accumulator, true
	}
	visited[line] = true
	defer func() { visited[line] = false }()
	inst := data[line]
	if so, ok := swappableOperations[inst.operation]; ok && !changed {
		altNextLine, newAcc := processInstruction(
			instruction{operation: so, argument: inst.argument}, accumulator,
			line,
		)
		if a, ok := process(data, visited, altNextLine, true, newAcc); ok {
			return a, true
		}
	}
	newLine, newAcc := processInstruction(inst, accumulator, line)
	return process(data, visited, newLine, changed, newAcc)
}

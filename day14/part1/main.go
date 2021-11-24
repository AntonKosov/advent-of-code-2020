package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type operation struct {
	address uint64
	value   uint64
}

type program struct {
	andMask    uint64
	orMask     uint64
	operations []operation
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([]program, error) {
	result := []program{}
	var pr *program
	handler := func(line string) error {
		if len(line) > 32 {
			if pr != nil {
				result = append(result, *pr)
			}
			pr = &program{
				andMask:    math.MaxUint64,
				orMask:     0,
				operations: []operation{},
			}
			var v uint64 = 1
			for i := 0; i < 36; i++ {
				switch rune(line[len(line)-i-1]) {
				case '0':
					pr.andMask &= ^v
				case '1':
					pr.orMask |= v
				}
				v <<= 1
			}
			return nil
		}
		split := strings.Split(line, " ")
		value, err := strconv.Atoi(split[2])
		if err != nil {
			return err
		}
		addressStr := split[0]
		address, err := strconv.Atoi(string(addressStr[4 : len(addressStr)-1]))
		if err != nil {
			return err
		}
		pr.operations = append(pr.operations, operation{
			address: uint64(address),
			value:   uint64(value),
		})

		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	result = append(result, *pr)
	return result, nil
}

func process(data []program) uint64 {
	mem := make(map[uint64]uint64)
	for _, pr := range data {
		for _, o := range pr.operations {
			mem[o.address] = (o.value & pr.andMask) | pr.orMask
		}
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

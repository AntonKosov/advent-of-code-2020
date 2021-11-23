package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

var (
	north aoc.Vector2
	east  aoc.Vector2
	south aoc.Vector2
	west  aoc.Vector2
)

func init() {
	north = aoc.Vector2{X: 0, Y: -1}
	east = north.RotateRight()
	south = east.RotateRight()
	west = south.RotateRight()
}

type instruction struct {
	action rune
	value  int
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
		value, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return err
		}
		result = append(result, instruction{
			action: rune(line[0]),
			value:  value,
		})
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

func process(data []instruction) int {
	orientation := east
	position := aoc.Vector2{}
	for _, i := range data {
		switch i.action {
		case 'N':
			position = position.Add(north.Mul(i.value))
		case 'S':
			position = position.Add(south.Mul(i.value))
		case 'E':
			position = position.Add(east.Mul(i.value))
		case 'W':
			position = position.Add(west.Mul(i.value))
		case 'L':
			for a := 0; a < i.value/90; a++ {
				orientation = orientation.RotateLeft()
			}
		case 'R':
			for a := 0; a < i.value/90; a++ {
				orientation = orientation.RotateRight()
			}
		case 'F':
			position = position.Add(orientation.Mul(i.value))
		default:
			panic("unknown action")
		}
	}
	return aoc.Abs(position.X) + aoc.Abs(position.Y)
}

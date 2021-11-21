package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type container map[string]int

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data, "shiny gold")
	fmt.Printf("Answer: %v\n", r)
}

func read() (map[string]container, error) {
	result := make(map[string]container)
	handler := func(line string) error {
		splitted := strings.Split(line, " bags contain ")
		bigBagColor := splitted[0]
		cont := container{}
		result[bigBagColor] = cont
		smallerBags := strings.Split(splitted[1], ", ")
		for _, smallerBag := range smallerBags {
			parts := strings.Split(smallerBag, " ")
			countStr := parts[0]
			if countStr == "no" {
				continue
			}
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return err
			}
			color := fmt.Sprintf("%v %v", parts[1], parts[2])
			cont[color] = count
		}

		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

func process(data map[string]container, color string) int {
	sum := 0
	for innerBagColor, count := range data[color] {
		sum += count * (1 + process(data, innerBagColor))
	}
	return sum
}

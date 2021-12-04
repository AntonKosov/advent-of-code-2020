package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type game struct {
	player1 []int
	player2 []int
}

func read() game {
	lines := aoc.ReadAllInput()
	g := game{}

	i := 1
	for ; ; i++ {
		line := lines[i]
		if line == "" {
			break
		}
		g.player1 = append(g.player1, aoc.StrToInt(line))
	}

	i += 2
	for ; ; i++ {
		line := lines[i]
		if line == "" {
			break
		}
		g.player2 = append(g.player2, aoc.StrToInt(line))
	}

	return g
}

func process(data game) int {
	var winner []int
	for len(winner) == 0 {
		p1 := data.player1[0]
		p2 := data.player2[0]
		if p1 > p2 {
			data.player1 = append(data.player1, p1, p2)
		} else {
			data.player2 = append(data.player2, p2, p1)
		}
		data.player1 = data.player1[1:]
		data.player2 = data.player2[1:]
		if len(data.player1) == 0 {
			winner = data.player2
		} else if len(data.player2) == 0 {
			winner = data.player1
		}
	}

	score := 0
	for i, c := range winner {
		score += c * (len(winner) - i)
	}

	return score
}

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

type state struct {
	player1 []int
	player2 []int
}

func read() state {
	lines := aoc.ReadAllInput()
	g := state{}

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

func process(data state) int {
	winner, _ := game(data.player1, data.player2, 1)

	score := 0
	for i, c := range winner {
		score += c * (len(winner) - i)
	}

	return score
}

func game(deck1, deck2 []int, gameN int) (w []int, number int) {
	for number == 0 {
		w, number = round(deck1, deck2, gameN)
	}

	return w, number
}

func round(deck1, deck2 []int, gameN int) (w []int, number int) {
	states := make(map[string]bool)
	for {
		state := fmt.Sprintf("%v|%v", deck1, deck2)
		if states[state] {
			return deck1, 1
		}
		states[state] = true

		c1 := deck1[0]
		deck1 = deck1[1:]
		c2 := deck2[0]
		deck2 = deck2[1:]

		var winner int
		if c1 <= len(deck1) && c2 <= len(deck2) {
			copy1, copy2 := copyDecks(deck1, c1, deck2, c2)
			_, winner = game(copy1, copy2, gameN+1)
		} else {
			if c1 > c2 {
				winner = 1
			} else if c1 < c2 {
				winner = 2
			} else {
				panic("Equal values")
			}
		}
		switch winner {
		case 1:
			deck1 = append(deck1, c1, c2)
			if len(deck2) == 0 {
				return deck1, 1
			}
		case 2:
			deck2 = append(deck2, c2, c1)
			if len(deck1) == 0 {
				return deck2, 2
			}
		default:
			panic("Unknown winner")
		}
	}
}

func copyDecks(d1 []int, count1 int, d2 []int, count2 int) (copy1, copy2 []int) {
	copy1 = make([]int, count1)
	copy2 = make([]int, count2)
	copy(copy1, d1)
	copy(copy2, d2)
	return copy1, copy2
}

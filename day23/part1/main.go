package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type game struct {
	labels  []*byte
	turns   int
	current byte
	buffer  []*byte
}

func (g *game) findIndex(v byte) *int {
	for i, l := range g.labels {
		if l != nil && v == *l {
			return &i
		}
	}
	return nil
}

func (g *game) moveToBuffer(index int) {
	for i := 0; i < 3; i++ {
		li := (index + i + 1) % 9
		g.buffer[i] = g.labels[li]
		g.labels[li] = nil
	}
}

func (g *game) returnFromBuffer(pi int) {
	pi++
	for {
		ni := (pi + 3) % 9
		g.labels[pi%9] = g.labels[ni]
		pi++
		if *g.labels[ni] == g.current {
			break
		}
	}
	for i := 0; i < 3; i++ {
		g.labels[(pi+i)%9] = g.buffer[i]
	}
}

func (g *game) makeTurn() {
	index := g.findIndex(g.current)
	g.moveToBuffer(*index)
	for {
		g.current--
		if g.current < 1 {
			g.current = 9
		}
		i := g.findIndex(g.current)
		if i == nil {
			continue
		}
		g.returnFromBuffer(*index)
		g.current = *g.labels[(*index+1)%9]
		break
	}
}

func (g *game) answer() string {
	var a strings.Builder
	index := *g.findIndex(1) + 1
	for {
		v := *g.labels[index%9]
		if v == 1 {
			return a.String()
		}
		a.WriteByte(v + '0')
		index++
	}
}

func read() game {
	lines := aoc.ReadAllInput()

	turns := aoc.StrToInt(lines[1])
	start := aoc.StrToInt(lines[2])
	g := game{
		turns:   turns,
		current: byte(start),
		buffer:  make([]*byte, 3),
	}

	for _, l := range lines[0] {
		v := byte(l) - byte('0')
		g.labels = append(g.labels, &v)
	}

	return g
}

func process(data game) string {
	for i := 0; i < data.turns; i++ {
		data.makeTurn()
	}
	return data.answer()
}

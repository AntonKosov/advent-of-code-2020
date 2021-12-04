package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

const maxLabel = 1_000_000

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type item struct {
	label int
	next  *item
}

type game struct {
	labels  map[int]*item
	turns   int
	current *item
	buffer  *item
}

func (g *game) moveToBuffer() {
	g.buffer = g.current.next
	lib := g.buffer.next.next
	g.current.next = lib.next
	lib.next = nil
}

func (g *game) returnFromBuffer(destination int) {
	di := g.labels[destination]
	n := di.next
	di.next = g.buffer
	g.buffer.next.next.next = n
	g.buffer = nil
}

func (g *game) isPresent(label int) bool {
	s := g.buffer
	for s != nil {
		if s.label == label {
			return false
		}
		s = s.next
	}
	return true
}

func (g *game) makeTurn() {
	g.moveToBuffer()
	destination := g.current.label
	for {
		destination--
		if destination < 1 {
			destination = maxLabel
		}
		if !g.isPresent(destination) {
			continue
		}
		g.returnFromBuffer(destination)
		g.current = g.current.next
		break
	}
}

func (g *game) answer() int64 {
	i1 := g.labels[1]
	return int64(i1.next.label) * int64(i1.next.next.label)
}

func read() game {
	lines := aoc.ReadAllInput()

	turns := aoc.StrToInt(lines[1])
	current := aoc.StrToInt(lines[2])
	g := game{
		turns:  turns,
		labels: make(map[int]*item, maxLabel),
	}

	var lastItem *item
	var firstItem *item
	addItem := func(label int) {
		newItem := &item{label: label}
		g.labels[label] = newItem
		if firstItem == nil {
			firstItem = newItem
		} else {
			lastItem.next = newItem
		}
		lastItem = newItem
	}
	for _, l := range lines[0] {
		label := int(byte(l) - byte('0'))
		addItem(label)
	}
	for l := 10; l <= maxLabel; l++ {
		addItem(l)
	}
	lastItem.next = firstItem
	g.current = g.labels[current]

	return g
}

func process(data game) int64 {
	for i := 0; i < data.turns; i++ {
		data.makeTurn()
	}
	return data.answer()
}

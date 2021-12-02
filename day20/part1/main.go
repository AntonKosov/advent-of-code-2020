package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

const pictureSize = 10

type tile struct {
	id             int
	availableEdges map[int64]bool
}

func newTile(id int, content []int64) tile {
	availableEdges := make(map[int64]bool, 8)

	edges := getEdges(content)
	for _, e := range edges {
		availableEdges[e] = true
		availableEdges[reverseBits(e)] = true
	}

	return tile{
		id:             id,
		availableEdges: availableEdges,
	}
}

func read() []tile {
	lines := aoc.ReadAllInput()
	var tiles []tile

	i := -1
	for i < len(lines)-1 {
		i++
		line := lines[i]
		if line == "" {
			continue
		}

		id := aoc.StrToInt(strings.Split(strings.Split(line, " ")[1], ":")[0])
		content := make([]int64, pictureSize)
		for j := 0; j < pictureSize; j++ {
			replaced0 := strings.ReplaceAll(lines[i+j+1], ".", "0")
			replacedAll := strings.ReplaceAll(replaced0, "#", "1")
			v, err := strconv.ParseInt(replacedAll, 2, 16)
			if err != nil {
				panic(err.Error())
			}
			content[j] = v
		}

		tiles = append(tiles, newTile(id, content))

		i += pictureSize
	}

	return tiles
}

func process(tiles []tile) int {
	neighbors := make(map[int]int)
	for _, tile := range tiles {
		for _, neighborTile := range tiles {
			if tile.id == neighborTile.id {
				continue
			}
			for edge := range tile.availableEdges {
				if neighborTile.availableEdges[edge] {
					neighbors[tile.id]++
					break
				}
			}
		}
	}

	result := 1
	count := 0
	for id, c := range neighbors {
		if c == 2 {
			result *= id
			count++
		}
	}

	if count != 4 {
		panic(fmt.Sprintf("Incorrect number of corners: %v", count))
	}

	return result
}

func reverseBits(v int64) int64 {
	var rv int64
	for i := 0; i < pictureSize; i++ {
		rv = (rv << 1) | ((v >> i) & 1)
	}
	return rv
}

func getEdges(pic []int64) []int64 {
	column := func(shift int) int64 {
		var v int64
		for i := 0; i < pictureSize; i++ {
			v = (v << 1) | ((pic[i] >> int64(shift)) & 1)
		}
		return v
	}

	result := []int64{pic[0], pic[pictureSize-1], column(0), column(pictureSize - 1)}

	return result
}

package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

const pictureSize = 10

type picture [][]rune

func newPicture(size int) picture {
	pic := make(picture, size)
	for i := range pic {
		pic[i] = make([]rune, size)
	}
	return pic
}

func (p picture) column(col int) string {
	var sb strings.Builder
	for i := 0; i < pictureSize; i++ {
		sb.WriteRune(p[col][i])
	}
	return sb.String()
}

func (p picture) row(row int) string {
	var sb strings.Builder
	for i := 0; i < pictureSize; i++ {
		sb.WriteRune(p[i][row])
	}
	return sb.String()
}

func (p picture) size() int {
	return len(p)
}

func (p picture) rotate() picture {
	r := newPicture(p.size())
	for x := 0; x < r.size(); x++ {
		for y := 0; y < r.size(); y++ {
			r[x][y] = p[y][p.size()-x-1]
		}
	}
	return r
}

func (p picture) flip() picture {
	f := newPicture(p.size())
	for x := 0; x < f.size(); x++ {
		for y := 0; y < f.size(); y++ {
			f[x][y] = p[x][p.size()-y-1]
		}
	}
	return f
}

func (p picture) countDots() int {
	sum := 0
	for _, c := range p {
		for _, d := range c {
			if d == '#' {
				sum++
			}
		}
	}
	return sum
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

type tile struct {
	edges map[direction]string
	pic   picture
}

func newTile(pic picture) tile {
	edges := map[direction]string{
		up:    pic.row(0),
		down:  pic.row(pictureSize - 1),
		left:  pic.column(0),
		right: pic.column(pictureSize - 1),
	}
	return tile{edges: edges, pic: pic}
}

type tileInfo struct {
	id            int
	chosenVariant *tile
	variants      []tile
}

func newTileInfo(id int, pic picture) tileInfo {
	var variants []tile

	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			variants = append(variants, newTile(pic))
			pic = pic.rotate()
		}
		pic = pic.flip()
	}

	return tileInfo{id: id, variants: variants}
}

func read() map[int]*tileInfo {
	lines := aoc.ReadAllInput()
	tiles := make(map[int]*tileInfo)

	i := -1
	for i < len(lines)-1 {
		i++
		line := lines[i]
		if line == "" {
			continue
		}

		id := aoc.StrToInt(strings.Split(strings.Split(line, " ")[1], ":")[0])
		pic := newPicture(pictureSize)
		for j := 0; j < pictureSize; j++ {
			pic[j] = []rune(lines[i+j+1])
		}

		t := newTileInfo(id, pic)
		tiles[id] = &t

		i += pictureSize
	}

	return tiles
}

func process(tiles map[int]*tileInfo) int {
	grid := buildGrid(tiles)
	pic := removeBorders(grid)

	return countDots(pic, []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	})
}

func removeBorders(grid [][]*tileInfo) picture {
	pieceSize := pictureSize - 2
	pic := newPicture(len(grid) * pieceSize)
	for x := 0; x < pic.size(); x++ {
		for y := 0; y < pic.size(); y++ {
			ti := grid[x/pieceSize][y/pieceSize]
			pic[x][y] = ti.chosenVariant.pic[1+x%pieceSize][1+y%pieceSize]
		}
	}
	return pic
}

func countDots(pic picture, pattern []string) int {
	pH := len(pattern)
	pW := len(pattern[0])
	var p []aoc.Vector2
	for y := 0; y < pH; y++ {
		for x := 0; x < pW; x++ {
			if pattern[y][x] == '#' {
				p = append(p, aoc.NewVector2(x, y))
			}
		}
	}

	for f := 0; f < 2; f++ {
		for i := 0; i < 4; i++ {
			sum := 0
			for x := 0; x < pic.size()-pW+1; x++ {
			next:
				for y := 0; y < pic.size()-pH+1; y++ {
					for _, v := range p {
						if pic[x+v.X][y+v.Y] != '#' {
							continue next
						}
					}
					sum++
				}
			}

			if sum > 0 {
				return pic.countDots() - sum*len(p)
			}
			pic = pic.rotate()
		}
		pic = pic.flip()
	}

	panic("Not found")
}

type tilesMap struct {
	tilesInfo map[int]*tileInfo
	sideSize  int
	found     map[int]bool
	grid      [][]*tileInfo
}

func (m *tilesMap) pickVariant(x, y int, tile *tileInfo, variant *tile) {
	tile.chosenVariant = variant
	m.grid[x][y] = tile
	m.found[tile.id] = true
}

func (m *tilesMap) resetVariant(x, y int, tile *tileInfo) {
	tile.chosenVariant = nil
	m.grid[x][y] = nil
	m.found[tile.id] = false
}

func fillMap(tilesMap *tilesMap, x, y int) bool {
	for tileID, tileInfo := range tilesMap.tilesInfo {
		if tilesMap.found[tileID] {
			continue
		}
		for _, v := range tileInfo.variants {
			if y > 0 && tilesMap.grid[x][y-1].chosenVariant.edges[down] != v.edges[up] {
				continue
			}
			if x > 0 && tilesMap.grid[x-1][y].chosenVariant.edges[right] != v.edges[left] {
				continue
			}
			tilesMap.pickVariant(x, y, tileInfo, &v)
			if x == tilesMap.sideSize-1 && y == tilesMap.sideSize-1 {
				return true
			}
			nx := x
			ny := y
			if x < tilesMap.sideSize-1 {
				nx++
			} else {
				nx = 0
				ny++
			}
			if fillMap(tilesMap, nx, ny) {
				return true
			}
			tilesMap.resetVariant(nx, ny, tileInfo)
		}
	}

	return false
}

func buildGrid(tiles map[int]*tileInfo) [][]*tileInfo {
	sideSize := int(math.Sqrt(float64(len(tiles))))
	tilesMap := tilesMap{
		tilesInfo: tiles,
		sideSize:  sideSize,
		found:     make(map[int]bool, len(tiles)),
		grid:      make([][]*tileInfo, sideSize),
	}
	for i := 0; i < sideSize; i++ {
		tilesMap.grid[i] = make([]*tileInfo, sideSize)
	}

	if !fillMap(&tilesMap, 0, 0) {
		panic("No solution")
	}

	return tilesMap.grid
}

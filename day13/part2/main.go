package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type bus struct {
	id     int
	offset int
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data []bus) {
	lines := aoc.ReadAllInput()
	buses := strings.Split(lines[0], ",")
	for i, b := range buses {
		if b == "x" {
			continue
		}
		id := aoc.StrToInt(b)
		data = append(data, bus{id: id, offset: i})
	}
	return data
}

type row struct {
	rem int
	mod int
}

// Chinese Reminder Theorem is used
func process(data []bus) int {
	var rows []row
	mb := 1
	for _, b := range data {
		rem := 0
		if b.offset > 0 {
			rem = b.id - b.offset
		}
		rows = append(rows, row{rem: rem, mod: b.id})
		mb *= b.id
	}

	sum := 0
	for _, r := range rows {
		base := mb / r.mod
		for i := 1; ; i++ {
			if (base*i)%r.mod == 1 {
				sum += r.rem * base * i
				break
			}
		}
	}

	return sum % mb
}

/*
// brute force :(
func process(data []bus) int64 {
	maxId := 0
	for i, b := range data {
		if data[maxId].id < b.id {
			maxId = i
		}
	}
	busWithMaxID := data[maxId]
	const chunkSize = 100000000
	semaphore := make(chan struct{}, runtime.NumCPU())
	stopSignal := make(chan struct{})
	stopCollectingSignal := make(chan struct{})
	results := []int64{}
	var collectingWG sync.WaitGroup
	var workersWG sync.WaitGroup
	result := make(chan int64)
	collectingWG.Add(1)
	go func() {
		defer collectingWG.Done()
		for {
			select {
			case <-stopCollectingSignal:
				return
			case t := <-result:
				if len(results) == 0 {
					close(stopSignal)
				}
				results = append(results, t)
			}
		}
	}()
	var m int64
loop:
	for m = 1; ; m += chunkSize {
		select {
		case <-stopSignal:
			break loop
		case semaphore <- struct{}{}:
			workersWG.Add(1)
			go func(first, last int64) {
				defer workersWG.Done()
				defer func() { <-semaphore }()
				var j int64
				var t int64
				for j = first; j <= last; j++ {
					t = busWithMaxID.id*j - busWithMaxID.offset
					found := true
					for _, b := range data {
						if (t+b.offset)%b.id != 0 {
							found = false
							break
						}
					}
					if found {
						result <- t
						return
					}
				}
				fmt.Printf("End processing: %v-%v, maxTimestamp: %v\n", first, last, t)
			}(m, m+chunkSize-1)
		}
	}

	workersWG.Wait()
	close(stopCollectingSignal)
	collectingWG.Wait()
	min := results[0]
	for i := 1; i < len(results); i++ {
		if results[i] < min {
			min = results[i]
		}
	}
	return min
}
*/

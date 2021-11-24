package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type bus struct {
	id     int64
	offset int64
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([]bus, error) {
	result := make([]bus, 0)
	handler := func(line string) error {
		buses := strings.Split(line, ",")
		for i, b := range buses {
			if b == "x" {
				continue
			}
			id, err := strconv.Atoi(b)
			if err != nil {
				return err
			}
			result = append(result, bus{
				id:     int64(id),
				offset: int64(i),
			})
		}
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	return result, nil
}

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
	// brute force :(
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

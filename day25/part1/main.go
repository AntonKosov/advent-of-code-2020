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

func read() []int {
	lines := aoc.ReadAllInput()

	return []int{
		aoc.StrToInt(lines[0]),
		aoc.StrToInt(lines[1]),
	}
}

func process(data []int) int {
	loopSize := solve(data[0])
	key := encryptionKey(loopSize, data[1])

	return key
}

const mod = 20201227

func encryptionKey(loopSize, subjectNumber int) int {
	ek := 1
	for i := 0; i < loopSize; i++ {
		ek = (ek * subjectNumber) % mod
	}
	return ek
}

func solve(publicKey int) (loopSize int) {
	const subjectNumber = 7
	for ek := 1; ek != publicKey; ek = (ek * subjectNumber) % mod {
		loopSize++
	}
	return loopSize
}

package aoc

import "strconv"

func VerifyIntBetween(numberSt string, min, max int) bool {
	number, err := strconv.Atoi(numberSt)
	if err != nil {
		return false
	}
	return number >= min && number <= max
}

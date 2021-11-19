package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	expenses, err := readExpenses()
	if err != nil {
		log.Fatalf("Error reading expenses: %v", err.Error())
	}
	r, err := processExpenses(expenses, 2020)
	if err != nil {
		log.Fatalf("Processing error: %v", err.Error())
	}
	fmt.Printf("Answer: %v\n", r)
}

func processExpenses(expenses []int, sum int) (int, error) {
	m := make(map[int]bool, len(expenses))
	for _, e := range expenses {
		v := sum - e
		if m[v] {
			return e * v, nil
		}
		m[e] = true
	}
	return 0, fmt.Errorf("not found")
}

func readExpenses() ([]int, error) {
	expenses := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exText := scanner.Text()
		ex, err := strconv.Atoi(exText)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, ex)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return expenses, nil
}

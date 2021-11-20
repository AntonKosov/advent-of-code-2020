package aoc

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(handler func(string) error) error {
	if len(os.Args) != 2 {
		return fmt.Errorf("wrong arguments")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := handler(line); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

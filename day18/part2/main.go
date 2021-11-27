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

func read() [][]string {
	lines := aoc.ReadAllInput()

	var data [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		data = append(data, strings.Split(line, " "))
	}

	return data
}

func process(data [][]string) int {
	sum := 0

	for _, ex := range data {
		sum += calcEx(ex)
	}

	return sum
}

func calcEx(ex []string) int {
	rpn := buildRPN(ex)
	return calcRPN(rpn)
}

func calcRPN(rpn []interface{}) int {
	var stack aoc.Stack
	for _, ins := range rpn {
		switch ins {
		case "*":
			n0 := stack.Pop().(int)
			n1 := stack.Pop().(int)
			stack.Push(n0 * n1)
		case "+":
			n0 := stack.Pop().(int)
			n1 := stack.Pop().(int)
			stack.Push(n0 + n1)
		default:
			stack.Push(ins)
		}
	}

	return stack.Pop().(int)
}

// buildRPN is based on https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func buildRPN(ex []string) []interface{} {
	var rpn []interface{}
	var opStack aoc.Stack

	precedence := map[string]int{"*": 1, "+": 2}

	for _, token := range ex {
		switch token {
		case "(":
			opStack.Push(token)
		case ")":
			for {
				st := opStack.Pop()
				if st == "(" {
					break
				}
				rpn = append(rpn, st)
			}
		case "+", "*":
			if !opStack.IsEmpty() {
				cp, cpOK := precedence[token]
				lp, lOK := precedence[opStack.Peek().(string)]
				if cpOK && lOK && cp <= lp {
					op := opStack.Pop()
					rpn = append(rpn, op)
				}
			}
			opStack.Push(token)
		default:
			rpn = append(rpn, aoc.StrToInt(token))
		}
	}

	for !opStack.IsEmpty() {
		rpn = append(rpn, opStack.Pop())
	}

	return rpn
}

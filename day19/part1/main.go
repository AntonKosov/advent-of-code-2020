package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

func main() {
	rules, messages := read()
	r := process(rules, messages)
	fmt.Printf("Answer: %v\n", r)
}

type matcher interface {
	verify(rules map[int]matcher, message string, start int) (last int, ok bool)
}

type runeMatcher struct {
	r rune
}

func (m runeMatcher) verify(_ map[int]matcher, message string, start int) (last int, ok bool) {
	return start, message[start] == byte(m.r)
}

type sequenceMatcher struct {
	sequence []int
}

func (m sequenceMatcher) verify(rules map[int]matcher, message string, start int) (last int, ok bool) {
	for _, rule := range m.sequence {
		matcher := rules[rule]
		last, ok = matcher.verify(rules, message, start)
		if !ok {
			return 0, false
		}
		start = last + 1
	}
	return last, true
}

type sequencesMatcher struct {
	sequences []sequenceMatcher
}

func (m sequencesMatcher) verify(rules map[int]matcher, message string, start int) (last int, ok bool) {
	for _, s := range m.sequences {
		last, ok = s.verify(rules, message, start)
		if ok {
			return last, true
		}
	}
	return 0, false
}

func read() (rules map[int]matcher, messages []string) {
	rules = make(map[int]matcher)
	lines := aoc.ReadAllInput()

	i := 0
	for {
		line := lines[i]
		if line == "" {
			break
		}

		sc := strings.Split(line, ":")
		ruleId := aoc.StrToInt(sc[0])
		var m matcher
		sp := sc[1]
		if strings.Contains(sp, "\"") {
			m = runeMatcher{r: rune(sp[2])}
		} else if strings.Contains(sp, "|") {
			sm := sequencesMatcher{}
			rs := strings.Split(sp, "|")
			for _, r := range rs {
				sm.sequences = append(sm.sequences, sequenceMatcher{sequence: aoc.StrToInts(r, " ")})
			}
			m = sm
		} else {
			m = sequenceMatcher{sequence: aoc.StrToInts(sp, " ")}
		}

		rules[ruleId] = m
		i++
	}

	for j := i + 1; j < len(lines); j++ {
		line := lines[j]
		if line != "" {
			messages = append(messages, line)
		}
	}

	return rules, messages
}

func process(rules map[int]matcher, messages []string) int {
	sum := 0

	for _, m := range messages {
		if last, ok := rules[0].verify(rules, m, 0); ok && last == len(m)-1 {
			sum++
		}
	}

	return sum
}

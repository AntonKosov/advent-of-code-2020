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
	for _, ruleID := range m.sequence {
		matcher := rules[ruleID]
		last, ok = matcher.verify(rules, message, start)
		if !ok {
			return 0, false
		}
		if last >= len(message)-1 {
			return last, true
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

func rule0(rules map[int]matcher, message string) bool {
	// must start with at least one 42 and has at least one pair 42 31
	// (there may be several pairs, e.g. 42 42 31 31)
	rule42 := rules[42]
	start := 0
	for {
		// rule 8
		last, ok := rule42.verify(rules, message, start)
		if !ok || last >= len(message)-1 {
			return false
		}
		start = last + 1
		if endsWithRule11(rules, message, start) {
			return true
		}
	}
}

func endsWithRule11(rules map[int]matcher, message string, start int) bool {
	rule42 := rules[42]

	// 42...42 31...31
	count := 0
	for {
		last, ok := rule42.verify(rules, message, start)
		if !ok || last >= len(message)-1 {
			return false
		}
		count++
		start = last + 1
		if endsWithRule31(rules, message, start, count) {
			return true
		}
	}
}

func endsWithRule31(rules map[int]matcher, message string, start, count int) bool {
	rule31 := rules[31]
	for i := 0; i < count; i++ {
		last, ok := rule31.verify(rules, message, start)
		if ok && i == count-1 && last == len(message)-1 {
			return true
		}
		if !ok || last >= len(message)-1 {
			return false
		}
		start = last + 1
	}
	return false
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
		if ok := rule0(rules, m); ok {
			sum++
		}
	}

	return sum
}

package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type ticket []int

type validValues map[int]bool

type tickets struct {
	titles        map[string]validValues
	myTicket      ticket
	nearbyTickets []ticket
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() tickets {
	lines := aoc.ReadAllInput()
	allValidValues := map[int]bool{}

	tickets := tickets{
		titles: map[string]validValues{},
	}

	addValidValues := func(title string, tokens ...string) {
		for _, t := range tokens {
			numbers := aoc.StrToInts(t, "-")
			for i := numbers[0]; i <= numbers[1]; i++ {
				allValidValues[i] = true
				vv, ok := tickets.titles[title]
				if !ok {
					vv = make(validValues)
					tickets.titles[title] = vv
				}
				vv[i] = true
			}
		}
	}

	i := 0
	for ; lines[i] != ""; i++ {
		line := lines[i]
		title := strings.Split(line, ":")[0]
		s := strings.Split(line, " ")
		addValidValues(title, s[len(s)-3], s[len(s)-1])
	}

	tickets.myTicket = append(tickets.myTicket, aoc.StrToInts(lines[i+2], ",")...)

	i += 5
nextLine:
	for ; lines[i] != ""; i++ {
		line := lines[i]
		var t ticket
		values := aoc.StrToInts(line, ",")
		for _, value := range values {
			if !allValidValues[value] {
				continue nextLine
			}
			t = append(t, value)
		}
		tickets.nearbyTickets = append(tickets.nearbyTickets, t)
	}

	return tickets
}

func process(data tickets) int {
	possibleColumns := map[string]map[int]bool{}
	for title, validValues := range data.titles {
	nextColumn:
		for column := 0; column < len(data.myTicket); column++ {
			for _, ticket := range data.nearbyTickets {
				if !validValues[ticket[column]] {
					continue nextColumn
				}
			}
			cis := possibleColumns[title]
			if cis == nil {
				cis = map[int]bool{}
				possibleColumns[title] = cis
			}
			cis[column] = true
		}
	}

	detectedColumns := map[string]int{}
	for i := 0; i < len(data.myTicket); i++ {
		for title, columns := range possibleColumns {
			if len(columns) != 1 {
				continue
			}
			var column int
			for key := range columns {
				column = key
				break
			}
			detectedColumns[title] = column
			delete(possibleColumns, title)
			for _, possibleColumn := range possibleColumns {
				delete(possibleColumn, column)
			}
			break
		}
	}

	mul := 1
	for title, column := range detectedColumns {
		if strings.HasPrefix(title, "departure") {
			mul *= data.myTicket[column]
		}
	}

	return mul
}

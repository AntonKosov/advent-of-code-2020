package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/AntonKosov/advent-of-code-2020/aoc"
)

type passport map[string]string

const cid = "cid"

var validator map[string]func(string) bool

func init() {
	validator = make(map[string]func(string) bool, 8)
	validator["byr"] = func(s string) bool {
		return aoc.VerifyIntBetween(s, 1920, 2002)
	}
	validator["iyr"] = func(s string) bool {
		return aoc.VerifyIntBetween(s, 2010, 2020)
	}
	validator["eyr"] = func(s string) bool {
		return aoc.VerifyIntBetween(s, 2020, 2030)
	}
	validator["hgt"] = func(s string) bool {
		hgt := string(s[:len(s)-2])
		if strings.HasSuffix(s, "in") {
			return aoc.VerifyIntBetween(hgt, 59, 76)
		}
		return aoc.VerifyIntBetween(hgt, 150, 193)
	}
	validator["hcl"] = func(s string) bool {
		matcher := regexp.MustCompile("^#[0-9a-f]{6}$")
		return matcher.Match([]byte(s))
	}
	validator["ecl"] = func(s string) bool {
		matcher := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
		return matcher.Match([]byte(s))
	}
	validator["pid"] = func(s string) bool {
		matcher := regexp.MustCompile("^[0-9]{9}$")
		return matcher.Match([]byte(s))
	}
	validator[cid] = func(_ string) bool { return true }
}

func main() {
	data, err := read()
	if err != nil {
		log.Fatalf("Error reading data: %v", err.Error())
	}
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func process(data []passport) int {
	valid := 0
next:
	for _, pass := range data {
		vf := len(pass)
		if _, ok := pass[cid]; ok {
			vf--
		}
		if vf != 7 {
			continue
		}
		for k, v := range pass {
			if !validator[k](v) {
				continue next
			}
		}
		valid++
	}

	return valid
}

func read() ([]passport, error) {
	result := []passport{}
	var pass map[string]string
	handler := func(line string) error {
		if line == "" {
			result = append(result, pass)
			pass = nil
			return nil
		}
		if pass == nil {
			pass = make(map[string]string, 8)
		}
		fields := strings.Split(line, " ")
		for _, f := range fields {
			kv := strings.Split(f, ":")
			pass[kv[0]] = kv[1]
		}
		return nil
	}
	if err := aoc.ReadInput(handler); err != nil {
		return nil, err
	}
	result = append(result, pass)
	return result, nil
}

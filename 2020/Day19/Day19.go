package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Char struct {
	char uint8
}

type Sequential struct {
	rules []*Rule
}

type Composite struct {
	rule1 *Rule
	rule2 *Rule
}

type Rule interface {
	Matches(str string) (bool, int)
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day19/input.txt")
	lines := strings.Split(string(bytes), "\n")

	rules, input := ParseRules(lines, func(s string) string { return s })
	println(CountMatches(lines[input:], rules[0]))

	rules, input = ParseRules(lines, func(s string) string {
		if strings.HasPrefix(s, "8: ") {
			return "8: 42 | 42 8"
		} else if strings.HasPrefix(s, "11: ") {
			return "11: 42 31 | 42 11 31"
		}
		return s
	})
	println(CountMatches(lines[input:], rules[0]))
}

func ParseRules(lines []string, transform func(string) string) ([]Rule, int) {
	// When the input section starts
	input := 0
	// Max rule #
	max := 0
	for i, line := range lines {
		if line == "" {
			input = i + 1
			break
		}
		split := strings.Split(line, ": ")
		n, _ := strconv.Atoi(split[0])

		// Expand array if necessary
		if n > max {
			max = n
		}
	}
	rules := make([]Rule, max+1)

	// Read Rules
	for _, line := range lines[:input-1] {
		line = transform(line)
		split := strings.Split(line, ": ")
		n, _ := strconv.Atoi(split[0])

		set := strings.Split(split[1], " | ")
		rule := Rule(nil)
		if len(set) == 2 {
			rule = &Composite{nil, nil}
		}
		for _, part := range set {
			refs := strings.Split(part, " ")
			r := Rule(nil)

			if len(refs) == 1 && strings.Contains(refs[0], "\"") {
				r = &Char{refs[0][1]}
			} else {
				nums := make([]int, len(refs))
				for j, s := range refs {
					nums[j], _ = strconv.Atoi(s)
				}
				all := make([]*Rule, len(nums))
				for j := range all {
					all[j] = &rules[nums[j]]
				}
				r = &Sequential{all}
			}

			if rule != nil {
				comp := rule.(*Composite)
				if comp.rule1 == nil {
					comp.rule1 = &r
				} else {
					comp.rule2 = &r
				}
			} else {
				rule = r
			}
		}
		rules[n] = rule
	}
	return rules, input
}

func CountMatches(lines []string, rule Rule) int {
	matches := 0
	for _, line := range lines {
		if match, _ := rule.Matches(line); match {
			matches++
		}
	}
	return matches
}

func (r *Char) Matches(str string) (bool, int) {
	if len(str) > 0 && str[0] == r.char {
		return true, 1
	} else {
		return false, 0
	}
}

func (r *Sequential) Matches(str string) (bool, int) {
	offset := 0
	for _, rule := range r.rules {
		m, c := (*rule).Matches(str[offset:])
		if !m {
			return false, 0
		}
		offset += c
	}
	return true, offset
}

func (r *Composite) Matches(str string) (bool, int) {
	if m1, c1 := (*r.rule1).Matches(str); m1 {
		return true, c1
	}
	if m2, c2 := (*r.rule2).Matches(str); m2 {
		return true, c2
	}
	return false, 0
}

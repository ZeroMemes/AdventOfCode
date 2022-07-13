package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Constants for Discovery state
const (
	FIELDS = iota // Reading all of the fields and their ruleset
	SELF          // Reading the ticket labeled "your ticket"
	AWAIT         // Waiting for the "nearby tickets" section
	NEARBY        // Reading the nearby tickets
)

type Candidates []bool

type Field struct {
	name string
	rule Rule
}

type Rule struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

type Ticket struct {
	data  []int
	valid bool
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day16/input.txt")
	lines := strings.Split(string(bytes), "\n")

	// All of the tickets, [0] = "Your Ticket"
	tickets := make([]*Ticket, 0)

	// All of the fields for numbers in the tickets
	fields := make([]*Field, 0)

	// Current discovery state
	state := FIELDS

	// RegExp for field definitions
	expr := regexp.MustCompile("([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")

	// Locate all field definitions and tickets
	for i, line := range lines {
		switch state {
		case FIELDS:
			if lines[i] == "" {
				state = SELF
				continue
			}
			match := expr.FindStringSubmatch(line)
			min1, _ := strconv.Atoi(match[2])
			max1, _ := strconv.Atoi(match[3])
			min2, _ := strconv.Atoi(match[4])
			max2, _ := strconv.Atoi(match[5])
			fields = append(fields, &Field{match[1], Rule{min1, max1, min2, max2}})
		case SELF:
			if line == "your ticket:" {
				tickets = append(tickets, ParseTicket(lines[i+1]))
				state = AWAIT
			}
		case AWAIT:
			if line == "nearby tickets:" {
				state = NEARBY
			}
		case NEARBY:
			tickets = append(tickets, ParseTicket(line))
		}
	}

	start := time.Now()
	err := 0
	for _, ticket := range tickets[1:] {
		hasErr := false
		for _, num := range ticket.data {
			// Whether or not this particular number was valid across any of the fields
			valid := false

			// Check all of the fields on this number
			for _, f := range fields {
				if f.rule.Accepts(num) {
					valid = true
				}
			}

			// If no field accepts the number, increment the "ticket scanning error rate"
			if !valid {
				err += num
				hasErr = true
			}
		}

		// If none of the fields had a problem matching at least 1 of the rules, the ticket is valid
		if !hasErr {
			ticket.valid = true
		}
	}
	fmt.Printf("%d (%dms)\n", err, time.Since(start).Milliseconds())
	start = time.Now()

	// Column Index -> Field Index -> Bool (Field candidacy for a given column)
	allCandidates := make([]Candidates, len(fields))
	for column := range allCandidates {
		allCandidates[column] = make(Candidates, len(fields))
		for field := range allCandidates[column] {
			allCandidates[column][field] = true
		}
	}

	// Field Index -> Column Index
	key := make([]int, len(fields))
	for i := range key {
		key[i] = -1 // Invalidate all fields
	}

	for (func() bool {
		// Continue looping while there's still >1 candidate for a given column
		for _, candidates := range allCandidates {
			if candidates.NRemaining() > 1 {
				return true
			}
		}
		return false
	})() {
		// Iterate through all of the columns and their field candidates
		for column, candidates := range allCandidates {
			// If this column has been resolved, set the index in the key and continue
			if rem := candidates.NRemaining(); rem == 1 {
				key[candidates.Remaining()[0]] = column
				continue
			}

			for field, viable := range candidates {
				// If we've already confirmed that this candidate isn't viable, ignore it
				if !viable {
					continue
				}
				// If this field has already been resolved, remove it as a candidate for this column
				if key[field] >= 0 {
					candidates[field] = false
					continue
				}
				// Check all of the remaining candidate fields against the tickets in this column
				for _, ticket := range tickets[1:] {
					if ticket.valid && !fields[field].rule.Accepts(ticket.data[column]) {
						candidates[field] = false
						break
					}
				}
			}
		}
	}

	// Find everything in our ticket starting with "departure" and get the product
	product := int64(1)
	for field, column := range key {
		if strings.HasPrefix(fields[field].name, "departure") {
			product *= int64(tickets[0].data[column])
		}
	}
	fmt.Printf("%d (%dms)\n", product, time.Since(start).Milliseconds())
}

func (r Rule) Accepts(val int) bool {
	return (val >= r.min1 && val <= r.max1) || (val >= r.min2 && val <= r.max2)
}

func (c Candidates) NRemaining() int {
	ret := 0
	for _, v := range c {
		if v {
			ret++
		}
	}
	return ret
}

func (c Candidates) Remaining() []int {
	ret := make([]int, 0)
	for i, v := range c {
		if v {
			ret = append(ret, i)
		}
	}
	return ret
}

func ParseTicket(str string) *Ticket {
	ret := make([]int, 0)
	for _, num := range strings.Split(str, ",") {
		i, err := strconv.Atoi(num)
		if err == nil {
			ret = append(ret, i)
		}
	}
	return &Ticket{ret, false}
}

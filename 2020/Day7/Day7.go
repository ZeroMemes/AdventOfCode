package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Bag struct {
	name string
	bags map[*Bag]int
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day7/input.txt")
	lines := strings.Split(string(bytes), "\n")

	// Allocate memory for all bags
	registry := make([]Bag, len(lines))

	// Create map of names to index and initialize all bags
	indices := map[string]int{}
	for i, s := range lines {
		name := strings.Split(s, " bag")[0]
		indices[name] = i
		registry[i] = NewBag(name)
	}

	// Function to return pointer to bag of name
	get := func(name string) *Bag {
		return &registry[indices[name]]
	}

	for _, line := range lines {
		split := strings.Split(strings.ReplaceAll(line, ".", ""), " contain ")
		name := strings.Split(split[0], " bag")[0]
		inner := strings.Split(split[1], ", ")

		bag := get(name)

		// Look at all inner bags
		for _, b := range inner {
			sep := strings.SplitN(b, " ", 2)
			name := strings.Split(sep[1], " bag")[0]
			count, _ := strconv.Atoi(sep[0])
			if name != "other" { // lol
				bag.bags[get(name)] = count
			}
		}
	}

	count := 0
	for _, bag := range registry {
		if bag.contains("shiny gold") {
			count++
		}
	}
	println(count)
	println(get("shiny gold").countTotalBags())
}

func NewBag(name string) Bag {
	return Bag{name, map[*Bag]int{}}
}

func (bag *Bag) contains(name string) bool {
	for b := range bag.bags {
		if b.name == name || b.contains(name) {
			return true
		}
	}
	return false
}

func (bag *Bag) countTotalBags() int {
	total := 0
	for b, count := range bag.bags {
		total += count + count*b.countTotalBags()
	}
	return total
}

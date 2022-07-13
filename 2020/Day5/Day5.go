package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day5/input.txt")
	split := strings.Split(string(bytes), "\n")

	lowest := math.MaxInt32
	highest := 0
	lol := make(map[int]bool)

	for _, seat := range split {
		s := seat[0:7]
		s = strings.ReplaceAll(s, "F", "0")
		s = strings.ReplaceAll(s, "B", "1")
		row, _ := strconv.ParseInt(s, 2, 0)

		s = seat[7:10]
		s = strings.ReplaceAll(s, "L", "0")
		s = strings.ReplaceAll(s, "R", "1")
		column, _ := strconv.ParseInt(s, 2, 0)

		id := int(row)*8 + int(column)
		if id > highest {
			highest = id
		}
		if id < lowest {
			lowest = id
		}
		lol[id] = true
	}

	println(highest)

	for i := lowest; i <= highest; i++ {
		if _, ok := lol[i]; !ok {
			println(i)
		}
	}
}

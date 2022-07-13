package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2021/Day2/input.txt")
	split := strings.Split(string(bytes), "\n")

	hpos, depth := 0, 0
	for _, s := range split {
		a, _ := strconv.Atoi(strings.Split(s, " ")[1])
		switch s[0] {
		case 'f':
			hpos += a
		case 'd':
			depth += a
		case 'u':
			depth -= a
		}
	}
	println(hpos * depth)

	hpos, depth, aim := 0, 0, 0
	for _, s := range split {
		a, _ := strconv.Atoi(strings.Split(s, " ")[1])
		switch s[0] {
		case 'f':
			hpos += a
			depth += a * aim
		case 'd':
			aim += a
		case 'u':
			aim -= a
		}
	}
	println(hpos * depth)
}

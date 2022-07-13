package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day2/input.txt")
	split := strings.Split(string(bytes), "\n")

	handle(split, func(low int, high int, char uint8, pass []uint8) bool {
		count := 0
		for _, c := range pass {
			if c == char {
				count++
			}
		}
		return count >= low && count <= high
	})

	handle(split, func(low int, high int, char uint8, pass []uint8) bool {
		return (pass[low-1] == char) != (pass[high-1] == char) // xor
	})
}

func handle(split []string, validate func(int, int, uint8, []uint8) bool) {
	valid := 0
	pattern := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")

	for _, str := range split {
		match := pattern.FindStringSubmatch(str)

		low, _ := strconv.Atoi(match[1])
		high, _ := strconv.Atoi(match[2])
		char := []uint8(match[3])[0]
		pass := []uint8(match[4])

		if validate(low, high, char, pass) {
			valid++
		}
	}

	println(valid)
}

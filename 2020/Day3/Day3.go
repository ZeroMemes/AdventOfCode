package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day3/input.txt")
	lines := strings.Split(string(bytes), "\n")

	println(trees(lines, 3, 1))
	println(
		trees(lines, 1, 1) *
			trees(lines, 3, 1) *
			trees(lines, 5, 1) *
			trees(lines, 7, 1) *
			trees(lines, 1, 2))
}

func trees(lines []string, right int, down int) int {
	width := len(lines[0])
	height := len(lines)

	x := 0
	count := 0

	for y := down; y < height; y += down {
		x += right
		x %= width

		if lines[y][x] == '#' {
			count++
		}
	}
	return count
}

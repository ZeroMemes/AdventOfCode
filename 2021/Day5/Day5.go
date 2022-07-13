package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type World map[Point]bool

func main() {
	bytes, _ := ioutil.ReadFile("2021/Day5/input.txt")
	split := strings.Split(string(bytes), "\n")

	lines := make([]Line, 0)

	world1 := World{}
	world2 := World{}

	pattern := regexp.MustCompile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

	for _, line := range split {
		match := pattern.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])
		lines = append(lines, Line{
			x1,
			y1,
			x2,
			y2,
		})
	}

	for _, line := range lines {
		dx := max(-1, min(1, line.x2-line.x1))
		dy := max(-1, min(1, line.y2-line.y1))
		p := Point{line.x1, line.y1}

		for {
			if line.IsStraight() {
				world1.Count(p)
			}
			world2.Count(p)

			if p.x == line.x2 && p.y == line.y2 {
				break
			}
			p.x += dx
			p.y += dy
		}
	}

	println(world1.CountOverlaps())
	println(world2.CountOverlaps())
}

func (w *World) Count(p Point) {
	_, ok := (*w)[p]
	(*w)[p] = ok
}

func (w *World) CountOverlaps() int {
	count := 0
	for coord := range *w {
		if duplicated, ok := (*w)[coord]; duplicated && ok {
			count++
		}
	}
	return count
}

func (l Line) IsStraight() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

package main

import (
	"io/ioutil"
	"strings"
)

const (
	SIZE = 99
)

type Forest struct {
	height [SIZE][SIZE]int
}

type Pos struct {
	x int
	y int
}

var UP = Pos{0, -1}
var DOWN = Pos{0, 1}
var LEFT = Pos{-1, 0}
var RIGHT = Pos{1, 0}

var CARDINAL = [4]Pos{
	UP, DOWN, LEFT, RIGHT,
}

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day8/input.txt")
	split := strings.Split(string(bytes), "\n")

	forest := Forest{}
	for y, row := range split {
		for x, ch := range row {
			forest.height[x][y] = int(ch - '0')
		}
	}
	println(forest.VisibleFromOutside())
	println(forest.MaxScenicScore())
}

func (f *Forest) VisibleFromOutside() int {
	visible := map[Pos]bool{}
	for d := 0; d < SIZE; d++ {
		scan := func(edge, offset Pos) {
			visible[edge] = true
			ref := f.HeightAt(edge)
			f.ScanFrom(edge, offset, func(pos Pos, h int) bool {
				if h > ref {
					visible[pos] = true
					ref = h
				}
				return true
			})
		}
		scan(Pos{d, 0}, DOWN)
		scan(Pos{SIZE - 1 - d, SIZE - 1}, UP)
		scan(Pos{0, d}, RIGHT)
		scan(Pos{SIZE - 1, SIZE - 1 - d}, LEFT)
	}
	return len(visible)
}

func (f *Forest) MaxScenicScore() int {
	max := 0
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			pos := Pos{x, y}
			ref := f.HeightAt(pos)
			dist := [4]int{}

			for i, delta := range CARDINAL {
				f.ScanFrom(pos, delta, func(pos Pos, h int) bool {
					if h <= ref {
						dist[i]++
						if h == ref {
							return false
						}
					}
					return true
				})
			}

			score := dist[0] * dist[1] * dist[2] * dist[3]
			if score > max {
				max = score
			}
		}
	}
	return max
}

func (f *Forest) HeightAt(p Pos) int {
	if f.InBounds(p) {
		return f.height[p.x][p.y]
	} else {
		panic("")
	}
}

func (f *Forest) InBounds(p Pos) bool {
	return p.x >= 0 && p.y >= 0 && p.x < SIZE && p.y < SIZE
}

func (f *Forest) ScanFrom(origin, step Pos, consumer func(Pos, int) bool) {
	p := Pos{origin.x + step.x, origin.y + step.y}
	for f.InBounds(p) && consumer(p, f.HeightAt(p)) {
		p = Pos{p.x + step.x, p.y + step.y}
	}
}

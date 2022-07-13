package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const (
	OCCUPIED = '#'
	EMPTY    = 'L'
	FLOOR    = '.'
)

var width int
var height int

type Vec2 struct {
	x int
	y int
}

var directions = [8]Vec2{
	{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1},
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day11/input.txt")
	rows := strings.Split(string(bytes), "\n")
	height = len(rows)
	width = len(rows[0])
	run(rows, 1)
	run(rows, 2)
}

func run(rows []string, part int) {
	start := time.Now()
	curr := 0
	last := 0
	for ok := true; ok; ok = curr != last {
		last = curr
		rows = update(rows, part)
		curr = getOccupied(rows)
	}
	fmt.Printf("%d (%dms)\n", curr, time.Since(start).Milliseconds())
}

func update(rows []string, part int) []string {
	// make a copy because we need to update the
	// data with reference to the previous data
	updated := make([]string, len(rows))
	copy(updated, rows)

	for y, row := range rows {
		for x, seat := range row {
			if seat != FLOOR {
				i := getSurroundingOccupied(rows, x, y, part)
				if i == 0 { // If an empty seat should become occupied
					seat = OCCUPIED
				} else if i >= part+3 { // If an occupied seat should become empty
					seat = EMPTY
				} else { // We're not going to be modifying the row, so continue
					continue
				}
				updated[y] = updated[y][:x] + string(seat) + updated[y][x+1:]
			}
		}
	}
	return updated
}

func getSurroundingOccupied(rows []string, x int, y int, part int) int {
	occupied := 0
	for _, d := range directions {
		switch part {
		case 1:
			if seat, _ := getSeat(rows, x+d.x, y+d.y); seat == OCCUPIED {
				occupied++
			}
		case 2:
			for i := 1; true; i++ {
				seat, exists := getSeat(rows, x+(d.x*i), y+(d.y*i))
				if !exists || seat != FLOOR {
					if seat == OCCUPIED {
						occupied++
					}
					break
				}
			}
		}
	}
	return occupied
}

func getSeat(rows []string, x int, y int) (uint8, bool) {
	if y >= 0 && x >= 0 && y < height && x < width {
		return rows[y][x], true
	}
	return FLOOR, false
}

func getOccupied(rows []string) int {
	return strings.Count(strings.Join(rows, ""), string(OCCUPIED))
}

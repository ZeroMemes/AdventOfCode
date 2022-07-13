package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Direction int8

const (
	EAST Direction = iota
	SOUTH
	WEST
	NORTH
)

type Pos struct {
	east  int
	north int
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day12/input.txt")
	lines := strings.Split(string(bytes), "\n")

	dir := EAST
	ship := Pos{0, 0}
	for _, line := range lines {
		move := line[0]
		val, _ := strconv.Atoi(line[1:])
		mvdir := Direction(-1)
		switch move {
		case 'L':
			dir -= Direction(val / 90)
			dir %= NORTH
		case 'R':
			dir += Direction(val / 90)
			dir %= NORTH
		case 'F':
			mvdir = dir
		default:
			mvdir = getDirection(move)
		}
		ship.move(mvdir, val)
	}
	println(ship.manhattan())

	ship = Pos{0, 0}
	wp := Pos{10, 1}
	for _, line := range lines {
		move := line[0]
		val, _ := strconv.Atoi(line[1:])
		switch move {
		case 'L':
			for i := 0; i < val/90; i++ {
				wp.east, wp.north = -wp.north, wp.east
			}
		case 'R':
			for i := 0; i < val/90; i++ {
				wp.east, wp.north = wp.north, -wp.east
			}
		case 'F':
			for i := 0; i < val; i++ {
				ship.east += wp.east
				ship.north += wp.north
			}
		default:
			wp.move(getDirection(move), val)
		}
	}
	println(ship.manhattan())
}

func getDirection(char uint8) Direction {
	return Direction(strings.IndexByte("ESWN", char))
}

func (p *Pos) move(dir Direction, val int) {
	switch dir {
	case NORTH:
		p.north += val
	case SOUTH:
		p.north -= val
	case EAST:
		p.east += val
	case WEST:
		p.east -= val
	}
}

func (p *Pos) manhattan() int {
	a, b := p.east, p.north
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

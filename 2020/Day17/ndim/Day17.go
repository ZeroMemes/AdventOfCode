package main

import (
	"io/ioutil"
	"math"
	"strings"
)

type Dimension struct {
	dimensions int
	grid       []*Position
}

type Position struct {
	vec    []int
	active bool
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day17/input.txt")
	lines := strings.Split(string(bytes), "\n")
	println(Run(3, lines))
	println(Run(4, lines))
}

func Run(dimensions int, lines []string) int {
	dim := NewDimension(dimensions)

	for x, line := range lines {
		for y, ch := range line {
			dim.GetPosition([]int{x, y}).active = ch == '#'
		}
	}

	for i := 0; i < 6; i++ {
		next := dim.Copy()
		for _, pos := range dim.grid {
			nearby := dim.GetSurrounding(pos)
			for _, p := range append(nearby, pos) {
				count := dim.GetSurroundingActive(p)
				if p.active {
					if count != 2 && count != 3 {
						next.GetPosition(p.vec).active = false
					}
				} else {
					if count == 3 {
						next.GetPosition(p.vec).active = true
					}
				}
			}
		}
		dim = next
	}
	count := 0
	for _, pos := range dim.grid {
		if pos.active {
			count++
		}
	}
	return count
}

func NewDimension(dimensions int) *Dimension {
	return &Dimension{dimensions, make([]*Position, 0)}
}

func (d *Dimension) Copy() *Dimension {
	dim := NewDimension(d.dimensions)
	for _, pos := range d.grid {
		dim.GetPosition(pos.vec).active = pos.active
	}
	return dim
}

func (d *Dimension) GetBounds() (min, max []int) {
	// It's fine that these gets initialized to all 0s
	// because it's guaranteed that 0,0,... exists
	min = make([]int, d.dimensions)
	max = make([]int, d.dimensions)
	for _, pos := range d.grid {
		for i, value := range pos.vec {
			if max[i] < value {
				max[i] = value
			} else if min[i] > value {
				min[i] = value
			}
		}
	}
	return
}

func (d *Dimension) GetPosition(vec []int) *Position {
	max := d.dimensions
	if len(vec) < max {
		max = len(vec)
	}
	matching := make([]int, d.dimensions)
	copy(matching, vec[:max])
	vec = matching

	// Find existing position with matching vec
	for _, pos := range d.grid {
		if SlicesMatch(pos.vec, vec) {
			return pos
		}
	}
	// Create the position
	pos := &Position{vec, false}
	d.grid = append(d.grid, pos)
	return pos
}

func (d *Dimension) GetSurrounding(pos *Position) []*Position {
	around := Pow(3, d.dimensions)
	offset := make([]int, d.dimensions)
	positions := make([]*Position, 0)
	for i := 0; i < around; i++ {
		// Calculate permutation of [-1, 0, 1] of length N dimensions
		val := i
		for j := 0; j < len(offset); j++ {
			offset[j] = (val % 3) - 1 + pos.vec[j]
			val /= 3
		}
		// Get Position at offset vec
		o := d.GetPosition(offset)
		if o != pos {
			positions = append(positions, o)
		}
	}
	return positions
}

func (d *Dimension) GetSurroundingActive(pos *Position) int {
	count := 0
	for _, o := range d.GetSurrounding(pos) {
		if o.active {
			count++
		}
	}
	return count
}

func (p *Position) Copy() *Position {
	vec := make([]int, len(p.vec))
	copy(vec, p.vec)
	return &Position{vec, p.active}
}

func SlicesMatch(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

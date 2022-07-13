package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

type Cache struct {
	get       []int
	translate []int
	size      []int
	offsets   []int
	flatten   []int
	surround  []int
}

type Dimension struct {
	dimensions int
	grid       []bool
	size       []int
	offsets    []int
	flatten    []int
	cache      Cache
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day17/input.txt")
	lines := strings.Split(string(bytes), "\n")
	Run(3, lines)
	Run(4, lines)
	Run(5, lines)
	Run(6, lines)
	Run(7, lines)
}

func Run(dimensions int, lines []string) {
	dim := NewDimension(dimensions)
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				pos := make([]int, dimensions)
				pos[0], pos[1] = x, y
				dim.SetActive(pos)
			}
		}
	}
	start := time.Now()
	for i := 0; i < 6; i++ {
		next := NewDimension(dimensions)
		dim.Expand(1)
		for index, active := range dim.grid {
			pos := dim.GetPos(index)
			count := dim.GetSurroundingActive(pos)
			if (active && (count == 2 || count == 3)) || (!active && count == 3) {
				next.SetActive(pos)
			}
		}
		dim = next
	}
	fmt.Printf("%d (%dms)\n", dim.GetActiveCount(), time.Since(start).Milliseconds())
}

func NewDimension(dimensions int) *Dimension {
	if dimensions < 2 {
		panic("Dimensions must be >= 2")
	}
	dim := &Dimension{
		dimensions,
		make([]bool, 1),
		make([]int, dimensions),
		make([]int, dimensions),
		make([]int, dimensions),
		Cache{
			make([]int, dimensions),
			make([]int, dimensions),
			make([]int, dimensions),
			make([]int, dimensions),
			make([]int, dimensions),
			make([]int, dimensions),
		},
	}
	ClearSlice(dim.size, 1)
	dim.CalculateFlatten(dim.size)
	copy(dim.flatten, dim.cache.flatten)
	return dim
}

func (d *Dimension) Accommodate(pos []int) {
	tr := d.Translate(pos)
	ClearSlice(d.cache.size, 0)
	ClearSlice(d.cache.offsets, 0)
	change := false
	for i, val := range tr {
		if val < 0 {
			d.cache.offsets[i] -= val
			d.cache.size[i] -= val
			change = true
		} else if val >= d.size[i] {
			d.cache.size[i] += val - d.size[i] + 1
			change = true
		}
	}
	if change {
		size := 1
		for i, val := range d.size {
			d.cache.size[i] += val
			size *= d.cache.size[i]
		}
		alloc := make([]bool, size)

		// Calculate the flatten values for new size and stores in cache.flatten
		d.CalculateFlatten(d.cache.size)

		// Translate offset cache by existing offsets
		for i, val := range d.offsets {
			d.cache.offsets[i] += val
		}

		// Migrate data
		for i := 0; i < len(d.grid); i++ {
			// Translate index to new grid
			n := d.GetIndex0(d.GetPos(i), d.cache.offsets, d.cache.flatten)
			alloc[n] = d.grid[i]
		}

		// Transfer to new grid
		d.grid = alloc
		copy(d.size, d.cache.size)
		copy(d.offsets, d.cache.offsets)
		copy(d.flatten, d.cache.flatten)
	}
}

func (d *Dimension) Expand(val int) {
	min := make([]int, d.dimensions)
	max := make([]int, d.dimensions)
	for i, offset := range d.offsets {
		min[i] = -offset - val
		max[i] = d.size[i] - offset - 1 + val
	}
	d.Accommodate(min)
	d.Accommodate(max)
}

func (d *Dimension) GetSurroundingActive(pos []int) int {
	count := 0
	around := Pow(3, d.dimensions)
	for i := 0; i < around; i++ {
		// Calculate permutation of [-1, 0, 1] of length N dimensions
		val := i
		ignore := true
		for j := 0; j < len(d.cache.surround); j++ {
			o := (val % 3) - 1
			if o != 0 {
				ignore = false
			}
			d.cache.surround[j] = o + pos[j]
			val /= 3
		}
		// Get Position at offset vec
		o := d.GetActive(d.cache.surround)
		if !ignore && o {
			count++
		}
	}
	return count
}

func (d *Dimension) SetActive(pos []int) {
	expand := false
	for i, v := range pos {
		v += d.offsets[i]
		if v >= d.size[i] || v < 0 {
			expand = true
			break
		}
	}
	if expand {
		d.Accommodate(pos)
	}
	index := d.GetIndex(pos)
	d.grid[index] = true
}

func (d *Dimension) GetActive(pos []int) bool {
	return d.GetActiveByIndex(d.GetIndex(pos))
}

func (d *Dimension) GetActiveByIndex(index int) bool {
	if index < 0 || index >= len(d.grid) {
		return false
	}
	return d.grid[index]
}

func (d *Dimension) GetIndex(pos []int) int {
	return d.GetIndex0(pos, d.offsets, d.flatten)
}

func (d *Dimension) GetIndex0(pos []int, offsets []int, flatten []int) int {
	index := 0
	for i, val := range d.Translate0(pos, offsets) {
		if val < 0 {
			return -1
		}
		index += val * flatten[i]
	}
	return index
}

func (d *Dimension) GetPos(index int) []int {
	return d.GetPos0(index, d.offsets, d.flatten)
}

func (d *Dimension) GetPos0(index int, offsets []int, flatten []int) []int {
	out := d.cache.get
	for i := range out {
		out[i] = index / flatten[i]
		if i != len(flatten)-1 {
			out[i] %= flatten[i+1] / flatten[i]
		}
	}
	for i, val := range offsets {
		out[i] -= val
	}
	return out
}

func (d *Dimension) Translate(pos []int) []int {
	return d.Translate0(pos, d.offsets)
}

func (d *Dimension) Translate0(pos []int, offsets []int) []int {
	for i, val := range pos {
		d.cache.translate[i] = val + offsets[i]
	}
	return d.cache.translate
}

func (d *Dimension) CalculateFlatten(size []int) []int {
	accum := 1
	for i, val := range size {
		d.cache.flatten[i] = accum
		accum *= val
	}
	return d.cache.flatten
}

func (d *Dimension) GetActiveCount() int {
	count := 0
	for _, active := range d.grid {
		if active {
			count++
		}
	}
	return count
}

func ClearSlice(slice []int, val int) {
	for i := range slice {
		slice[i] = val
	}
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

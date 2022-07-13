package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type D1 map[int]D2
type D2 map[int]D3
type D3 map[int]D4
type D4 map[int]bool

type Dimension struct {
	grid  D1
	use4D bool
}

type Vec4 struct {
	x int
	y int
	z int
	w int
}

type AABB struct {
	min Vec4
	max Vec4
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day17/input.txt")
	lines := strings.Split(string(bytes), "\n")
	Run(NewDimension(false, lines), 6)
	Run(NewDimension(true, lines), 6)
}

func Run(dim *Dimension, iter int) {
	start := time.Now()
	for i := 0; i < iter; i++ {
		next := dim.Copy()
		dim.IterateArea(dim.GetBounds().expand(1), func(pos Vec4) {
			around := dim.GetActiveNeighbors(pos)
			if dim.IsActive(pos) {
				if around != 2 && around != 3 {
					next.SetActive(pos, false)
				}
			} else {
				if around == 3 {
					next.SetActive(pos, true)
				}
			}
		})
		dim = next
	}
	fmt.Printf("%d (%dms)\n", dim.GetActiveCount(), time.Since(start).Milliseconds())
}

func NewDimension(use4D bool, lines []string) *Dimension {
	dim := &Dimension{make(D1), use4D}
	if lines != nil {
		for x, line := range lines {
			for y, ch := range line {
				dim.SetActive(Vec4{x, y, 0, 0}, ch == '#')
			}
		}
	}
	return dim
}

func NewAABB() AABB {
	return AABB{Vec4{0, 0, 0, 0}, Vec4{0, 0, 0, 0}}
}

func (dim *Dimension) GetBounds() AABB {
	bounds := NewAABB()
	for x, d2 := range dim.grid {
		for y, d3 := range d2 {
			for z, d4 := range d3 {
				for w := range d4 {
					bounds.Accommodate(x, y, z, w)
				}
			}
		}
	}
	return bounds
}

func (dim *Dimension) IterateArea(bounds AABB, callback func(pos Vec4)) {
	if !dim.use4D {
		bounds.min.w = 0
		bounds.max.w = 0
	}
	for x := bounds.min.x; x <= bounds.max.x; x++ {
		for y := bounds.min.y; y <= bounds.max.y; y++ {
			for z := bounds.min.z; z <= bounds.max.z; z++ {
				for w := bounds.min.w; w <= bounds.max.w; w++ {
					callback(Vec4{x, y, z, w})
				}
			}
		}
	}
}

func (dim *Dimension) ForEachLocation(callback func(pos Vec4)) {
	dim.IterateArea(dim.GetBounds(), callback)
}

func (dim *Dimension) GetActiveCount() int {
	count := 0
	dim.ForEachLocation(func(pos Vec4) {
		if dim.IsActive(pos) {
			count++
		}
	})
	return count
}

func (dim *Dimension) GetActiveNeighbors(pos Vec4) int {
	count := 0
	dim.IterateArea(AABB{pos.Sub(1), pos.Add(1)}, func(p Vec4) {
		if p != pos && dim.IsActive(p) {
			count++
		}
	})
	return count
}

func (dim *Dimension) Copy() *Dimension {
	ret := NewDimension(dim.use4D, nil)
	for x, plane := range dim.grid {
		ret.grid[x] = make(D2)
		for y, line := range plane {
			ret.grid[x][y] = make(D3)
			for z, fourth := range line {
				ret.grid[x][y][z] = make(D4)
				for w, active := range fourth {
					ret.grid[x][y][z][w] = active
				}
			}
		}
	}
	return ret
}

func (dim *Dimension) SetActive(pos Vec4, active bool) {
	x, y, z, w, d := pos.x, pos.y, pos.z, pos.w, dim.grid
	if _, ok := d[x]; !ok {
		d[x] = make(D2)
	}
	if _, ok := d[x][y]; !ok {
		d[x][y] = make(D3)
	}
	if _, ok := d[x][y][z]; !ok {
		d[x][y][z] = make(D4)
	}
	d[x][y][z][w] = active
}

func (dim *Dimension) IsActive(pos Vec4) bool {
	x, y, z, w, d := pos.x, pos.y, pos.z, pos.w, dim.grid
	if _, ok := d[x]; !ok {
		d[x] = make(D2)
	}
	if _, ok := d[x][y]; !ok {
		d[x][y] = make(D3)
	}
	if _, ok := d[x][y][z]; !ok {
		d[x][y][z] = make(D4)
	}
	if _, ok := d[x][y][z]; !ok {
		d[x][y][z][w] = false
	}
	return d[x][y][z][w]
}

func (a *AABB) Accommodate(x, y, z, w int) {
	if a.max.x < x {
		a.max.x = x
	} else if a.min.x > x {
		a.min.x = x
	}
	if a.max.y < y {
		a.max.y = y
	} else if a.min.y > y {
		a.min.y = y
	}
	if a.max.z < z {
		a.max.z = z
	} else if a.min.z > z {
		a.min.z = z
	}
	if a.max.w < w {
		a.max.w = w
	} else if a.min.w > w {
		a.min.w = w
	}
}

func (a AABB) expand(val int) AABB {
	return AABB{
		a.min.Sub(val),
		a.max.Add(val),
	}
}

func (v Vec4) Add(val int) Vec4 {
	return Vec4{
		v.x + val,
		v.y + val,
		v.z + val,
		v.w + val,
	}
}

func (v Vec4) Sub(val int) Vec4 {
	return v.Add(-val)
}

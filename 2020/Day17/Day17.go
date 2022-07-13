package main

import (
	"io/ioutil"
	"strings"
)

type Dimension map[int]Plane
type Plane map[int]Line
type Line map[int]bool
type Vec3 struct {
	x int
	y int
	z int
}
type AABB struct {
	min Vec3
	max Vec3
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day17/input.txt")
	lines := strings.Split(string(bytes), "\n")

	dimension := make(Dimension)
	dimension.initializePlane(lines)

	for i := 0; i < 6; i++ {
		newState := dimension.copy()
		dimension.iterateArea(dimension.getBounds().expand(1), func(x, y, z int) {
			active := dimension.isActive(x, y, z)
			around := dimension.getActiveNeighbors(x, y, z)
			if active {
				if around != 2 && around != 3 {
					newState.setActive(x, y, z, false)
				}
			} else {
				if around == 3 {
					newState.setActive(x, y, z, true)
				}
			}
		})
		dimension = newState
	}
	println(dimension.countActive())
}

func (d Dimension) getBounds() AABB {
	bounds := AABB{Vec3{0, 0, 0}, Vec3{0, 0, 0}}
	for x, plane := range d {
		if bounds.max.x < x {
			bounds.max.x = x
		}
		if bounds.min.x > x {
			bounds.min.x = x
		}
		for y, line := range plane {
			if bounds.max.y < y {
				bounds.max.y = y
			}
			if bounds.min.y > y {
				bounds.min.y = y
			}
			for z := range line {
				if bounds.max.z < z {
					bounds.max.z = z
				}
				if bounds.min.z > z {
					bounds.min.z = z
				}
			}
		}
	}
	return bounds
}

func (a AABB) expand(val int) AABB {
	return AABB{
		Vec3{
			a.min.x - val,
			a.min.y - val,
			a.min.z - val,
		},
		Vec3{
			a.max.x + val,
			a.max.y + val,
			a.max.z + val,
		},
	}
}

func (d Dimension) iterateArea(bounds AABB, callback func(x, y, z int)) {
	for x := bounds.min.x; x <= bounds.max.x; x++ {
		for y := bounds.min.y; y <= bounds.max.y; y++ {
			for z := bounds.min.z; z <= bounds.max.z; z++ {
				callback(x, y, z)
			}
		}
	}
}

func (d Dimension) forEachLocation(callback func(x, y, z int)) {
	d.iterateArea(d.getBounds(), callback)
}

func (d Dimension) countActive() int {
	count := 0
	d.forEachLocation(func(x, y, z int) {
		if d.isActive(x, y, z) {
			count++
		}
	})
	return count
}

func (d Dimension) initializePlane(lines []string) {
	for x, line := range lines {
		for y, ch := range line {
			d.setActive(x, y, 0, ch == '#')
		}
	}
}

func (d Dimension) getActiveNeighbors(x, y, z int) int {
	count := 0
	for ox := -1; ox <= 1; ox++ {
		for oy := -1; oy <= 1; oy++ {
			for oz := -1; oz <= 1; oz++ {
				if ox != 0 || oy != 0 || oz != 0 {
					if d.isActive(x+ox, y+oy, z+oz) {
						count++
					}
				}
			}
		}
	}
	return count
}

func (d Dimension) copy() Dimension {
	ret := make(Dimension)
	for x, plane := range d {
		ret[x] = make(Plane)
		for y, line := range plane {
			ret[x][y] = make(Line)
			for z, active := range line {
				ret[x][y][z] = active
			}
		}
	}
	return ret
}

func (d Dimension) setActive(x, y, z int, active bool) {
	if _, ok := d[x]; !ok {
		d[x] = make(Plane)
	}
	if _, ok := d[x][y]; !ok {
		d[x][y] = make(Line)
	}
	d[x][y][z] = active
}

func (d Dimension) isActive(x, y, z int) bool {
	if _, ok := d[x]; !ok {
		d[x] = make(Plane)
	}
	if _, ok := d[x][y]; !ok {
		d[x][y] = make(Line)
	}
	if _, ok := d[x][y][z]; !ok {
		d[x][y][z] = false
	}
	return d[x][y][z]
}

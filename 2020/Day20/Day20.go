package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Side int8

const (
	TOP Side = iota
	BOTTOM
	LEFT
	RIGHT
)

type Tile struct {
	id      int
	index   int
	data    [10][10]bool
	matches map[Side][]*Tile
}

const (
	ImageSize  = 12
	SeaMonster = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `
)

type Assembly [ImageSize][ImageSize]*Tile
type TileUsage [ImageSize * ImageSize]bool
type Image [ImageSize * 8][ImageSize * 8]bool

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day20/input.txt")
	lines := strings.Split(string(bytes), "\n")

	header := regexp.MustCompile("Tile ([0-9]{4}):")
	tiles := make([]*Tile, 0)

	start := time.Now()

	for i, line := range lines {
		match := header.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		id, _ := strconv.Atoi(match[1])
		tile := MakeTile(id, len(tiles))
		for j, data := range lines[i+1 : i+11] {
			for k, ch := range data {
				tile.data[j][k] = ch == '#'
			}
		}
		tiles = append(tiles, tile)
	}

	// lol
	if len(tiles) != ImageSize*ImageSize {
		panic("Input does not match for defined ImageSize")
	}

	// Keeping track of transformations? I prefer creating them automatically instead
	for _, tile := range tiles {
		for i := 1; i <= 0b1111; i++ {
			flipX := i&0b0001 > 0
			flipY := i&0b0010 > 0
			rotation := i & 0b1100 >> 2
			transformed := MakeTile(tile.id, tile.index)

			// Handle flip
			for row := range tile.data {
				for col, val := range tile.data[row] {
					j, k := row, col
					if flipX {
						j = 9 - j
					}
					if flipY {
						k = 9 - k
					}
					transformed.data[j][k] = val
				}
			}

			// Handle rotate
			for j := 0; j < rotation; j++ {
				original := transformed.data
				for row := range original {
					for col, val := range original[row] {
						// Rotate 90 degrees
						transformed.data[col][9-row] = val
					}
				}
			}

			tiles = append(tiles, transformed)
		}
	}

	// Index all of the matched sides
	for _, a := range tiles {
		for _, b := range tiles {
			if a.id == b.id {
				continue
			}
			for side := TOP; side <= RIGHT; side++ {
				if a.GetSide(side) == b.GetSide(side.Opposite()) {
					a.matches[side] = append(a.matches[side], b)
				}
			}
		}
	}
	fmt.Printf("Finished setup (%dms)\n", time.Since(start).Milliseconds())

	// Part 1
	start = time.Now()
	assembled := Assembly{}
	for _, tile := range tiles {
		// Setup State
		used := TileUsage{}
		used[tile.index] = true
		attempts := 0

		// Search
		if config, found := FindConfiguration(0, 0, used, Assembly{{tile}}, &attempts); found {
			assembled = config
			fmt.Printf("Part 1: %d (%dms)\n", assembled.MultiplyCorners(), time.Since(start).Milliseconds())
			break
		}
	}

	// Part 2
	start = time.Now()
	for _, image := range assembled.StitchAndTransform() {
		if monsters, found := image.FindSeaMonsters(); found {
			roughness := 0
			for y := range image {
				for x, val := range image[y] {
					if !monsters[y][x] && val {
						roughness++
					}
				}
			}
			fmt.Printf("Part 2: %d (%dms)\n", roughness, time.Since(start).Milliseconds())
			break
		}
	}
}

func MakeTile(id int, index int) *Tile {
	return &Tile{id, index, [10][10]bool{}, make(map[Side][]*Tile)}
}

func FindConfiguration(x int, y int, used TileUsage, out Assembly, attempts *int) (Assembly, bool) {
	next := ([]*Tile)(nil)

	// We'll be able to find the correct configuration in ImageSize^2 attempts, a direct route. Disregard anything else
	// (I actually don't know if this would be the case for a specially crafted input)
	if *attempts >= ImageSize*ImageSize {
		return out, false
	}
	*attempts++

	// Determine next tile position and find the next tile candidates
	x++
	if x >= ImageSize { // Next line
		y++
		x = 0
		if y >= ImageSize {
			return out, true
		} else {
			next = out[y-1][x].matches[BOTTOM]
		}
	} else { // To the right of current position
		if y > 0 {
			next = GetCommonTiles(out[y][x-1].matches[RIGHT], out[y-1][x].matches[BOTTOM])
		} else {
			next = out[y][x-1].matches[RIGHT]
		}
	}

	for _, n := range next {
		if !used[n.index] {
			// Copy state
			usedCopy := used
			outCopy := out

			// Modify
			outCopy[y][x] = n
			usedCopy[n.index] = true

			if o, found := FindConfiguration(x, y, usedCopy, outCopy, attempts); found {
				return o, true
			}
		}
	}

	return out, false
}

func GetCommonTiles(m1 []*Tile, m2 []*Tile) []*Tile {
	ret := make([]*Tile, 0)
	for _, t1 := range m1 {
		for _, t2 := range m2 {
			if t1 == t2 {
				ret = append(ret, t1)
			}
		}
	}
	return ret
}

func (i Image) FindSeaMonsters() (image Image, found bool) {
	monster := strings.Split(SeaMonster, "\n")
	for y, row := range i[:len(i)-3] {
	outer:
		for x := range row[:len(row)-20] {
			for my, mr := range monster {
				for mx, m := range mr {
					if m == '#' && !i[y+my][x+mx] {
						continue outer
					}
				}
			}
			found = true
			for my, mr := range monster {
				for mx, m := range mr {
					if m == '#' {
						image[y+my][x+mx] = true
					}
				}
			}
		}
	}
	return
}

func (a Assembly) StitchAndTransform() [16]Image {
	images := [16]Image{}
	for col := range a {
		for row, tile := range a[col] {
			for dy, dr := range tile.data[1:9] {
				for dx, val := range dr[1:9] {
					images[0][col*8+dy][row*8+dx] = val
				}
			}
		}
	}

	ref := images[0]

	for i := 1; i <= 0b1111; i++ {
		flipX := i&0b0001 > 0
		flipY := i&0b0010 > 0
		rotation := i & 0b1100 >> 2
		transformed := ref

		// Handle flip
		for row := range ref {
			for col, val := range ref[row] {
				j, k := row, col
				if flipX {
					j = (ImageSize*8 - 1) - j
				}
				if flipY {
					k = (ImageSize*8 - 1) - k
				}
				transformed[j][k] = val
			}
		}

		// Handle rotate
		for j := 0; j < rotation; j++ {
			original := transformed
			for row := range original {
				for col, val := range original[row] {
					// Rotate 90 degrees
					transformed[col][(ImageSize*8-1)-row] = val
				}
			}
		}
		images[i] = transformed
	}

	return images
}

func (a Assembly) MultiplyCorners() int {
	return a[0][0].id *
		a[0][ImageSize-1].id *
		a[ImageSize-1][0].id *
		a[ImageSize-1][ImageSize-1].id
}

func (s Side) Opposite() Side {
	switch s {
	case TOP:
		return BOTTOM
	case BOTTOM:
		return TOP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	}
	panic("Bad Side Provided!")
}

func (t *Tile) GetSide(side Side) [10]bool {
	switch side {
	case TOP:
		return t.data[0]
	case BOTTOM:
		return t.data[9]
	case LEFT:
		array := [10]bool{}
		for row := range t.data {
			array[row] = t.data[row][0]
		}
		return array
	case RIGHT:
		array := [10]bool{}
		for row := range t.data {
			array[row] = t.data[row][9]
		}
		return array
	}
	panic("Bad Side Provided!")
}

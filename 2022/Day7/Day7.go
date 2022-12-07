package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	parent  *dir
	files   map[string]int
	subdirs map[string]*dir
	size    int
}

func mkdir(parent *dir) *dir {
	return &dir{
		parent,
		make(map[string]int),
		make(map[string]*dir),
		0,
	}
}

func (d *dir) mkdir(name string) *dir {
	if lookup, ok := d.subdirs[name]; ok {
		return lookup
	} else {
		nd := mkdir(d)
		d.subdirs[name] = nd
		return nd
	}
}

func (d *dir) file(name string, size int) {
	d.files[name] = size
}

func (d *dir) walk(sizes *[]int) {
	for _, child := range d.files {
		d.size += child
	}
	for _, child := range d.subdirs {
		child.walk(sizes)
		d.size += child.size
	}
	*sizes = append(*sizes, d.size)
}

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day7/input.txt")
	split := strings.Split(string(bytes), "\n")

	cd := mkdir(nil)
	root := cd

	for _, s := range split {
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "$":
			if parts[1] == "cd" {
				if parts[2] == ".." {
					cd = cd.parent
				} else if parts[2] != "/" {
					cd = cd.mkdir(parts[2])
				}
			} else if parts[1] == "ls" {
				// do nothing?
			}
		case "dir":
			cd.mkdir(parts[1])
		default:
			size, _ := strconv.Atoi(parts[0])
			cd.file(parts[1], size)
		}
	}

	sizes := make([]int, 0)
	root.walk(&sizes)

	total := 0
	for _, val := range sizes {
		if val <= 100000 {
			total += val
		}
	}
	println(total)

	sort.Ints(sizes)
	for _, val := range sizes {
		if 70000000-root.size+val > 30000000 {
			println(val)
			return
		}
	}
}

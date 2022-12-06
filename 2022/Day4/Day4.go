package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day4/input.txt")
	split := strings.Split(string(bytes), "\n")

	p1 := 0
	p2 := 0
	for _, str := range split {
		pairs := strings.Split(str, ",")
		a := strings.Split(pairs[0], "-")
		b := strings.Split(pairs[1], "-")
		amin, _ := strconv.Atoi(a[0])
		amax, _ := strconv.Atoi(a[1])
		bmin, _ := strconv.Atoi(b[0])
		bmax, _ := strconv.Atoi(b[1])

		if (amin <= bmin && bmax <= amax) || (bmin <= amin && amax <= bmax) {
			p1++
		}
		if (bmax >= amin && bmax <= amax) || (amax >= bmin && amax <= bmax) ||
			(bmin >= amin && bmin <= amax) || (amin >= bmin && amin <= bmax) {
			p2++
		}
	}
	println(p1)
	println(p2)
}

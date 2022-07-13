package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2021/Day6/input.txt")

	ages := map[int]int{}
	for i := 0; i <= 8; i++ {
		ages[i] = 0
	}

	for _, str := range strings.Split(string(bytes), ",") {
		a, _ := strconv.Atoi(str)
		ages[a]++
	}

	for day := 0; day < 256; day++ {
		reset := ages[0]
		ages[0] = 0
		for i := 0; i <= 8; i++ {
			ages[i] = ages[i+1]
		}
		ages[8] += reset
		ages[6] += reset
	}

	total := 0
	for _, num := range ages {
		total += num
	}
	println(total)
}

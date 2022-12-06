package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day1/input.txt")
	split := strings.Split(string(bytes), "\n\n")

	totals := make([]int, len(split))
	for i, str := range split {
		nums := strings.Split(str, "\n")
		sum := 0
		for _, n := range nums {
			i, _ := strconv.Atoi(n)
			sum += i
		}
		totals[i] = sum
	}
	sort.Ints(totals)
	println(totals[len(totals)-1])
	println(totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3])
}

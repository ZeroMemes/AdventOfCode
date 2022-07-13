package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var cache = map[int]int64{}
var jolts []int

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day10/input.txt")
	lines := strings.Split(string(bytes), "\n")
	jolts = make([]int, len(lines))

	for i, str := range lines {
		num, _ := strconv.ParseInt(str, 10, 64)
		jolts[i] = int(num)
	}
	sort.Ints(jolts)

	// Start at 0 and end at +3 of largest
	jolts = append([]int{0}, append(jolts, jolts[len(jolts)-1]+3)...)

	diff := map[int]int{}
	for i := range jolts[1:] {
		diff[jolts[i+1]-jolts[i]]++
	}
	println(diff[1] * diff[3])
	println(count(0))
}

func count(n int) int64 {
	if val, ok := cache[n]; ok {
		return val
	}

	var ret int64
	if n >= jolts[len(jolts)-1] {
		ret = 1
	} else {
		for _, v := range jolts {
			if v > n && v <= n+3 {
				ret += count(v)
			}
		}
	}
	cache[n] = ret
	return ret
}

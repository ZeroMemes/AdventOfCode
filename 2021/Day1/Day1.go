package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2021/Day1/input.txt")
	split := strings.Split(string(bytes), "\n")
	count := len(split)

	nums := make([]int, count)
	for i, s := range split {
		nums[i], _ = strconv.Atoi(s)
	}

	println(countIncrement(nums, 1))
	println(countIncrement(nums, 3))
}

func countIncrement(nums []int, window int) int {
	prev, inc := 0, 0
	for i := 0; i < len(nums)-(window-1); i++ {
		sum := 0
		for j := i; j < i+window; j++ {
			sum += nums[j]
		}

		if prev > 0 && sum > prev {
			inc++
		}
		prev = sum
	}
	return inc
}

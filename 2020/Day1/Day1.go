package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day1/input.txt")
	split := strings.Split(string(bytes), "\n")
	count := len(split)

	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i], _ = strconv.Atoi(split[i])
	}

	handleA(nums)
	handleB(nums)
}

func handleA(nums []int) {
	count := len(nums)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if nums[i]+nums[j] == 2020 {
				println(nums[i] * nums[j])
				return
			}
		}
	}
}

func handleB(nums []int) {
	count := len(nums)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			for k := 0; k < count; k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					println(nums[i] * nums[j] * nums[k])
					return
				}
			}
		}
	}
}

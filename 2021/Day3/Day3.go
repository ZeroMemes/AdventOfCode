package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2021/Day3/input.txt")
	split := strings.Split(string(bytes), "\n")

	ones := map[int]int{}
	zeros := map[int]int{}

	for i := 0; i < 12; i++ {
		ones[i] = 0
		zeros[i] = 0
	}

	for _, s := range split {
		for j, c := range s {
			switch c {
			case '1':
				ones[j] = ones[j] + 1
			case '0':
				zeros[j] = zeros[j] + 1
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < 12; i++ {
		if ones[i] > zeros[i] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 0)
	e, _ := strconv.ParseInt(epsilon, 2, 0)
	println(g * e)

	nums := map[string]bool{}
	for _, s := range split {
		nums[s] = true
	}

	for position := 0; position < 12; position++ {
		ones := map[string]bool{}
		zeros := map[string]bool{}

		for s := range nums {
			switch s[position] {
			case '1':
				ones[s] = true
			case '0':
				zeros[s] = true
			}
		}

		if len(zeros) <= len(ones) {
			for s := range ones {
				delete(nums, s)
			}
		} else {
			for s := range zeros {
				delete(nums, s)
			}
		}

		if len(nums) == 1 {
			fmt.Println(nums)
			break
		}
	}
}

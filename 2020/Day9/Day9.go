package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day9/input.txt")
	lines := strings.Split(string(bytes), "\n")
	nums := make([]int64, len(lines))

	for i, str := range lines {
		num, _ := strconv.ParseInt(str, 10, 64)
		nums[i] = num
	}

	p1 := int64(0)
outer1:
	for i, num := range nums[25:] {
		for j, n1 := range nums[i : i+25] {
			for k, n2 := range nums[i : i+25] {
				if j != k && n1+n2 == num {
					continue outer1
				}
			}
		}
		p1 = num
		break
	}
	println(p1)

outer2:
	for i, num := range nums {
		sum := num
		smallest := int64(math.MaxInt64)
		largest := int64(0)
		for _, next := range nums[i+1:] {
			if next < smallest {
				smallest = next
			}
			if next > largest {
				largest = next
			}
			sum += next
			if sum > p1 {
				continue outer2
			} else if sum == p1 {
				println(largest + smallest)
				break outer2
			}
		}
	}
}

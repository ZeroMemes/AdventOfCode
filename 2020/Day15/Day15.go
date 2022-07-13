package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day15/input.txt")
	line := strings.Split(string(bytes), "\n")[0]
	split := strings.Split(line, ",")

	spoken := make(map[int][]int)
	nums := make([]int, len(split))
	for i, str := range split {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}

	start := time.Now()
	diff := 0
	for i := 1; true; i++ {
		num := 0

		if i > len(nums) {
			num = diff
		} else {
			num = nums[i-1]
		}

		data, ok := spoken[num]
		if !ok {
			data = []int{i, i}
			spoken[num] = data
		} else {
			data[0], data[1] = data[1], i
		}
		diff = data[1] - data[0]

		if i == 2020 {
			fmt.Printf("%d (%dms)\n", num, time.Since(start).Milliseconds())
		}
		if i == 30000000 {
			fmt.Printf("%d (%dms)\n", num, time.Since(start).Milliseconds())
			break
		}
	}
}

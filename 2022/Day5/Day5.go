package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day5/input.txt")
	split := strings.Split(string(bytes), "\n")

	pt1 := make([]string, 9)
	pt2 := make([]string, 9)

	for i, str := range split {
		if i < 8 {
			for j := 0; j < 9; j++ {
				ch := rune(str[j*4+1])
				if ch != ' ' {
					pt1[j] = pt1[j] + string(ch)
					pt2[j] = pt2[j] + string(ch)
				}
			}
		} else if i > 9 {
			parts := strings.Split(str, " ")
			cnt, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])

			from--
			to--

			for j := 0; j < cnt; j++ {
				rem := pt1[from][:1]
				pt1[from] = pt1[from][1:]
				pt1[to] = rem + pt1[to]
			}

			rem := pt2[from][0:cnt]
			pt2[from] = pt2[from][cnt:]
			pt2[to] = rem + pt2[to]
		}
	}
	for _, s := range pt1 {
		fmt.Print(s[:1])
	}
	fmt.Println()
	for _, s := range pt2 {
		fmt.Print(s[:1])
	}
}

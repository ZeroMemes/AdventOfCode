package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day13/input.txt")
	lines := strings.Split(string(bytes), "\n")
	earliest, _ := strconv.Atoi(lines[0])
	buses := map[int]int{}

	for i, s := range strings.Split(lines[1], ",") {
		v, err := strconv.Atoi(s)
		if err == nil {
			buses[v] = i
		}
	}
outer1:
	for i := earliest; true; i++ {
		for k := range buses {
			if i%k == 0 {
				println((i - earliest) * k)
				break outer1
			}
		}
	}

	// Solved using Wolfram Alpha :(
	for k, v := range buses {
		fmt.Printf("0=((x + %d) mod %d), ", v, k)
	}
}

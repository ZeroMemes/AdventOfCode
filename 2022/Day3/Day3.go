package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day3/input.txt")
	split := strings.Split(string(bytes), "\n")
	part1(split)
	part2(split)
}

func part1(split []string) {
	sum := 0
	for _, str := range split {
		first := str[:len(str)/2]
		second := str[len(str)/2:]

	outer:
		for _, a := range first {
			for _, b := range second {
				if a == b {
					if a >= 'A' && a <= 'Z' {
						sum += int(a-'A') + 27
					} else {
						sum += int(a-'a') + 1
					}
					break outer
				}
			}
		}
	}
	println(sum)
}

func part2(split []string) {
	sum := 0
	for i := 0; i < len(split)-2; i += 3 {
	outer2:
		for _, a := range split[i] {
			for _, b := range split[i+1] {
				for _, c := range split[i+2] {
					if a == b && b == c {
						if a >= 'A' && a <= 'Z' {
							sum += int(a-'A') + 27
						} else {
							sum += int(a-'a') + 1
						}
						break outer2
					}
				}
			}
		}
	}
	println(sum)
}

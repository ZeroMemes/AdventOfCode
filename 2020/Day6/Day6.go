package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day6/input.txt")
	split := strings.Split(string(bytes), "\n\n")

	count := 0
	all := 0
	for _, s := range split {
		m := make(map[uint8]int)
		people := strings.Split(s, "\n")
		s = strings.Join(people, "")
		for i := range s {
			m[s[i]]++
		}
		count += len(m)
		for _, v := range m {
			if v == len(people) {
				all++
			}
		}
	}
	println(count)
	println(all)
}

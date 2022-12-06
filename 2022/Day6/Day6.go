package main

import (
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("2022/Day6/input.txt")
	input := string(bytes)
	println(findUniqueSubstring(input, 4))
	println(findUniqueSubstring(input, 14))
}

func findUniqueSubstring(str string, length int) int {
	buf := ""
	for i, ch := range str {
		buf += string(ch)
		if len(buf) > length {
			buf = buf[1:]
		}
		if len(buf) == length && hasUniqueChars(buf) {
			return i + 1
		}
	}
	panic("")
}

func hasUniqueChars(str string) bool {
	m := make(map[rune]bool)
	for _, c := range str {
		if _, ok := m[c]; !ok {
			m[c] = true
		} else {
			return false
		}
	}
	return true
}

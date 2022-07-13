package main

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day14/input.txt")
	lines := strings.Split(string(bytes), "\n")
	execute(lines, part1)
	execute(lines, part2)
}

func execute(lines []string, writer func(string, map[int64]int64, int64, int64)) {
	mask := ""
	mem := make(map[int64]int64)
	expr := regexp.MustCompile("mem\\[([0-9]+)] = ([0-9]+)")

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, " = ")[1]
		} else {
			matches := expr.FindStringSubmatch(line)
			addr, _ := strconv.ParseInt(matches[1], 10, 64)
			val, _ := strconv.ParseInt(matches[2], 10, 64)
			writer(mask, mem, val, addr)
		}
	}

	sum := int64(0)
	for _, v := range mem {
		sum += v
	}
	println(sum)
}

func part1(mask string, mem map[int64]int64, val int64, addr int64) {
	for i, c := range mask {
		if c != 'X' {
			bit, _ := strconv.Atoi(string(c))
			if bit == 1 {
				val |= 1 << (35 - i)
			} else {
				val &^= 1 << (35 - i)
			}
		}
	}
	mem[addr] = val
}

func part2(mask string, mem map[int64]int64, val int64, raw int64) {
	addr := strconv.FormatInt(raw, 2)
	addr = strings.Repeat("0", 36-len(addr)) + addr

	floating := 0
	for i, c := range mask {
		if c != '0' {
			addr = addr[:i] + string(mask[i]) + addr[i+1:]
		}
		if c == 'X' {
			floating++
		}
	}

	// Max occurences is 9, so this is doable
	for i := 0; i < pow(2, floating); i++ {
		addVal := int64(0)
		x := 0
		for j, c := range addr {
			bit := 0
			if c == 'X' {
				bit = i >> (floating - x - 1) & 1
				x++
			} else if c == '1' {
				bit = 1
			}
			addVal |= int64(bit) << int64(35-j)
		}
		mem[addVal] = val
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Board [5][5]int
type Marked [5][5]bool

func main() {
	bytes, _ := ioutil.ReadFile("2021/Day4/input.txt")
	split := strings.Split(string(bytes), "\n")
	count := len(split)

	numstr := strings.Split(split[0], ",")
	nums := make([]int, len(numstr))
	for i, s := range numstr {
		nums[i], _ = strconv.Atoi(s)
	}

	boards := map[*Board]*Marked{}
	won := map[*Board]bool{}
	for i := 2; i < count; i += 6 {
		board := Board{}
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				num, _ := strconv.Atoi(strings.TrimSpace(split[i+row][col*3 : col*3+2]))
				board[row][col] = num
			}
		}
		boards[&board] = &Marked{}
		won[&board] = false
	}

	for _, num := range nums {
		for board, marked := range boards {
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					if board[row][col] == num {
						marked[row][col] = true
					}
				}
			}
		}

		for board, marked := range boards {
			if marked.HasSolution() && !won[board] {
				sum := 0
				for row := 0; row < 5; row++ {
					for col := 0; col < 5; col++ {
						if !marked[row][col] {
							sum += board[row][col]
						}
					}
				}
				println(sum * num)
				won[board] = true
			}
		}
	}
}

func (m *Marked) HasSolution() bool {
	for row := 0; row < 5; row++ {
		if m[row][0] && m[row][1] && m[row][2] && m[row][3] && m[row][4] {
			return true
		}
	}
	for col := 0; col < 5; col++ {
		if m[0][col] && m[1][col] && m[2][col] && m[3][col] && m[4][col] {
			return true
		}
	}
	return false
}

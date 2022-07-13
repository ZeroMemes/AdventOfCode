package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type OpHandler func(left int, right int, op uint8, operators string) (int, bool)

var operandExpr = regexp.MustCompile("[*+]")
var operatorExpr = regexp.MustCompile("[^*+]")

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day18/input.txt")
	lines := strings.Split(strings.ReplaceAll(string(bytes), " ", ""), "\n")

	// Part 1
	println(sum(lines, func(left int, right int, op uint8, operators string) (int, bool) {
		switch op {
		case '+':
			return left + right, true
		case '*':
			return left * right, true
		}
		return 0, false
	}))

	// Part 2
	println(sum(lines, func(left int, right int, op uint8, operators string) (int, bool) {
		switch op {
		case '+':
			return left + right, true
		case '*':
			if !strings.ContainsRune(operators, '+') {
				return left * right, true
			}
		}
		return 0, false
	}))
}

func sum(lines []string, handler OpHandler) int {
	ret := 0
	for _, line := range lines {
		ret += eval(line, handler)
	}
	return ret
}

func eval(str string, handler OpHandler) int {
	// Evaluate anything in parenthesis first, we didn't have
	// to do it for the part 1 but we have to do it now
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			cnt := 1
			end := 0
			for j, chj := range str[i+1:] {
				switch chj {
				case '(':
					cnt++
				case ')':
					cnt--
				}
				if cnt == 0 {
					end = j + i + 1
					break
				}
			}
			return eval(str[:i]+strconv.Itoa(eval(str[i+1:end], handler))+str[end+1:], handler)
		}
	}

	operators := operatorExpr.ReplaceAllString(str, "")
	split := operandExpr.Split(str, -1)
	operands := make([]int, len(split))
	for i, s := range split {
		operands[i], _ = strconv.Atoi(s)
	}

	for len(operands) > 1 {
		for i := range operators {
			// Pass to handle function
			val, success := handler(operands[i], operands[i+1], operators[i], operators)

			// If it succeeded, consume the right operand and the operator used
			if success {
				operands[i] = val
				operands = append(operands[:i+1], operands[i+2:]...)
				operators = operators[:i] + operators[i+1:]
				// Break because we modified the string we're iterating over
				break
			}
			continue
		}
	}
	return operands[0]
}

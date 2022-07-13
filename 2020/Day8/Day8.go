package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type OpCode uint8
type InsnList []Insn

const (
	NOP OpCode = iota
	JMP
	ACC
)

type Insn struct {
	opcode  OpCode
	operand int
}

type Interpreter struct {
	insns       InsnList
	executed    []bool
	accumulator int
	counter     int
	mutations   int
}

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day8/input.txt")
	lines := strings.Split(string(bytes), "\n")

	// Read program
	insns := parseProgram(lines)

	// Setup initial interpreter
	interpreters := list.List{}
	interpreters.PushBack(&Interpreter{insns, make([]bool, len(insns)), 0, 0, 0})

	start := time.Now()

	// While at least 1 interpreter is running...
	for interpreters.Len() > 0 {
		keep := list.List{}
		for e := interpreters.Front(); e != nil; e = e.Next() {
			i := (e.Value).(*Interpreter)

			// Solution to part 1
			if i.mutations == 0 && i.executed[i.counter] {
				fmt.Printf("1: %d (%dms)\n", i.accumulator, time.Since(start).Milliseconds())
				continue
			}

			// Solution to part 2
			if i.mutations == 1 && i.counter == len(i.insns) {
				fmt.Printf("2: %d (%dms)\n", i.accumulator, time.Since(start).Milliseconds())
				continue
			}

			// Remove anything out of the scope of what we're testing, or would've have already validated
			if i.mutations > 1 || i.counter >= len(i.insns) || i.executed[i.counter] {
				continue
			}

			// Handle mutation
			insn := i.insns[i.counter]
			if insn.opcode < ACC {
				branch := i.copy()
				keep.PushBack(branch)

				// Mutate accordingly
				if insn.opcode == NOP {
					branch.insns[i.counter].opcode = JMP
				} else {
					branch.insns[i.counter].opcode = NOP
				}
				branch.mutations++
			}

			// Once we've created a mutation we're free to update and keep the interpreter
			i.tick()
			keep.PushBack(i)
		}
		interpreters = keep
	}
}

func (i InsnList) copy() InsnList {
	c := make(InsnList, len(i))
	copy(c, i)
	return c
}

func (i *Interpreter) copy() *Interpreter {
	executed := make([]bool, len(i.insns))
	copy(executed, i.executed)
	return &Interpreter{
		i.insns.copy(),
		executed,
		i.accumulator,
		i.counter,
		i.mutations,
	}
}

func (i *Interpreter) tick() {
	insn := i.insns[i.counter]
	i.executed[i.counter] = true

	switch insn.opcode {
	case NOP:
		i.counter++
	case JMP:
		i.counter += insn.operand
	case ACC:
		i.accumulator += insn.operand
		i.counter++
	}
}

func parseProgram(lines []string) InsnList {
	insns := make(InsnList, len(lines))

	// Load Program
	for i, line := range lines {
		split := strings.Split(line, " ")
		value, _ := strconv.Atoi(split[1])

		var opcode OpCode
		switch split[0] {
		case "nop":
			opcode = NOP
		case "jmp":
			opcode = JMP
		case "acc":
			opcode = ACC
		}

		insns[i] = Insn{opcode, value}
	}

	return insns
}

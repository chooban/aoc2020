package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chooban/aoc2020/internal/io"
)

type instruction struct {
	command  string
	argument string
}

func main() {
	lines := io.ReadFileLines(os.Args[1])
	instructions := make([]instruction, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		instructions[i] = instruction{
			parts[0], parts[1],
		}
	}

	for idx, inst := range instructions {
		fmt.Println("Executing instruction", idx)
		if inst.command == "jmp" {
			newInstructions := make([]instruction, len(instructions))
			copy(newInstructions, instructions)
			newInstructions[idx] = instruction{
				"nop",
				instructions[idx].argument,
			}
			ret, acc := execute(newInstructions)
			if ret == 0 {
				fmt.Println("A zero!")
				fmt.Println(acc)
				os.Exit(0)
			}
		} else if inst.command == "nop" {
			newInstructions := make([]instruction, len(instructions))
			copy(newInstructions, instructions)
			newInstructions[idx] = instruction{
				"jmp",
				instructions[idx].argument,
			}
			ret, acc := execute(newInstructions)
			if ret == 0 {
				fmt.Println("A zero!")
				fmt.Println(acc)
				os.Exit(0)
			}
		} else {
			fmt.Println("Not changing", instructions[idx])
		}
	}
}

func execute(instructions []instruction) (int, int) {
	ret := 0
	acc := 0
	idx := 0
	executed := make([]bool, len(instructions))
	for {
		if idx >= len(instructions) {
			break
		}
		if executed[idx] == true {
			fmt.Printf("Instruction %d already executed\n", idx)
			ret = 1
			break
		}
		ins := instructions[idx]
		executed[idx] = true

		// fmt.Println("Executing", ins.command, ins.argument)
		switch ins.command {
		case "nop":
			idx++
		case "acc":
			val, err := strconv.Atoi(ins.argument)
			if err != nil {
				panic(1)
			}
			acc += val
			idx++
		case "jmp":
			val, err := strconv.Atoi(ins.argument)
			if err != nil {
				panic(1)
			}
			idx += val
		}
	}

	return ret, acc
}

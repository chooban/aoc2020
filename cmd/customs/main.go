package main

import (
	"fmt"
	"math/bits"
	"os"

	"github.com/chooban/aoc2020/internal/io"
)

type Key int

func main() {
	lines := io.ReadFileLines(os.Args[1])
	groups := [][]string{{}}

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			groups[len(groups)-1] = append(groups[len(groups)-1], lines[i])
		} else {
			if i != len(lines)-1 {
				groups = append(groups, []string{})
			}
		}
	}

	fmt.Println(len(groups), "groups to process")

	count := 0
	for _, group := range groups {
		uniqAnswers := map[string]bool{}
		for _, member := range group {
			for _, answer := range member {
				uniqAnswers[string(answer)] = true
			}
		}
		count += len(uniqAnswers)
	}

	fmt.Println(count, "yeses in total")

	part2Yeses := 0
	for _, group := range groups {
		groupAnswers := uint(0)
		for idx, member := range group {
			var memberAnswers = uint(0)

			for _, answer := range member {
				memberAnswers = setBit(
					memberAnswers,
					uint(int(answer)-96),
				)
			}

			if idx == 0 {
				groupAnswers = memberAnswers
			} else {
				groupAnswers = groupAnswers & memberAnswers
			}
		}
		part2Yeses += bits.OnesCount(uint(groupAnswers))
	}

	fmt.Println(part2Yeses, "correct yeses")
}

func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return n
}

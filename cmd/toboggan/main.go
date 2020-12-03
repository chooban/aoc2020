package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chooban/aoc2020/internal/toboggan"
)

func main() {

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	patterns := [][]int{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}

	trees := 1
	for _, values := range patterns {
		moveRight := values[0]
		moveDown := values[1]

		localTrees := toboggan.HowManyTrees(lines, moveRight, moveDown)
		trees *= localTrees

		fmt.Printf(
			"Encountered %d trees for %d, %d\n",
			localTrees,
			moveRight, moveDown,
		)
	}

	fmt.Printf("Total trees: %d", trees)
}

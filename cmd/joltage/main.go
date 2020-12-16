package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/chooban/aoc2020/internal/io"
)

func main() {
	lines := io.ReadFileLines(os.Args[1])

	ratings := make([]int, len(lines))
	// ratings[0] = 0

	for i, val := range lines {
		if len(val) == 0 {
			continue
		}
		nval, _ := strconv.Atoi(val)
		ratings[i+1] = nval
	}

	sort.Ints(ratings)
	ratings = append(ratings, ratings[len(ratings)-1]+3)

	fmt.Println("Joltage ratings:", ratings)

	part1(ratings)
	part2(ratings)
}

func part1(ratings []int) {
	ones := 0
	threes := 0

	for i := 1; i < len(ratings); i++ {
		previous := ratings[i-1]
		current := ratings[i]

		switch current - previous {
		case 1:
			ones++
		case 3:
			threes++
		}
	}

	fmt.Println("Ones:", ones, "Threes:", threes)
	fmt.Println("Product:", ones*threes)
}

func part2(ratings []int) {

	fmt.Println(ratings)
	adapters := map[int]bool{}
	for _, a := range ratings {
		adapters[a] = true
	}

	paths := map[int]int{}
	paths[0] = 1

	for _, rating := range ratings {
		for i := 1; i <= 3; i++ {
			nextRating := rating + i

			if adapters[nextRating] == true {
				paths[nextRating] += paths[rating]
			}
		}
	}

	fmt.Println(paths)
	fmt.Println(paths[ratings[len(ratings)-1]])
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/chooban/aoc2020/internal/io"
)

var window = 25

func main() {
	lines := io.ReadFileLines(os.Args[1])
	allNumbers := make([]int, len(lines)-1)

	for i, val := range lines {
		if len(val) == 0 {
			continue
		}
		nval, _ := strconv.Atoi(val)
		allNumbers[i] = nval
	}

	lookingFor := findInvalidSum(allNumbers, window)
	fmt.Println("No match for", lookingFor)

	head, tail := findContiguousRange(allNumbers, lookingFor)

	fmt.Println("Head:", head, "Tail:", tail)
	slice := allNumbers[head : tail+1]
	fmt.Println("Sum of slice", sumArray(slice))
	sort.Ints(slice)

	fmt.Println(slice)
	fmt.Println(slice[0] + slice[len(slice)-1])
}

func sumArray(n []int) (c int) {
	c = 0
	for _, v := range n {
		c += v
	}
	return
}

func findContiguousRange(allNumbers []int, target int) (head, tail int) {
	head = 0
	tail = 1
	sum := allNumbers[head] + allNumbers[tail]

	for {
		if sum == target {
			break
		} else if sum < target {
			tail++
			sum += allNumbers[tail]
		} else if sum > target {
			sum -= allNumbers[head]
			head++
		}
	}
	return
}

func findInvalidSum(allNumbers []int, window int) int {
	for i := window; i < len(allNumbers); i++ {
		nSlice := allNumbers[i-window : i]

		// Make a map to do the quick check
		values := map[int]bool{}
		for _, v := range nSlice {
			values[v] = true
		}

		found := false
		for k := range values {
			if allNumbers[i] == k {
				continue
			}
			if values[allNumbers[i]-k] == true {
				found = true
			}
		}
		if !found {
			return allNumbers[i]
		}
	}
	return 0
}

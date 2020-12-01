package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/chooban/aoc2020/internal/expenses"
)

func main() {
	fmt.Println("Reading", os.Args[1])
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	numbers := []int{}

	for _, i := range lines {
		j, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		numbers = append(numbers, j)
	}

	fmt.Println("Double answer is", expenses.FindSumToPair(numbers, 2020))
	fmt.Println("Product is", Product(expenses.FindSumToPair(numbers, 2020)))
	fmt.Println("Triple answer is", expenses.FindSumToTriple(numbers, 2020))
	fmt.Println("Product is", Product(expenses.FindSumToTriple(numbers, 2020)))
}

func Product(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}
	return result
}

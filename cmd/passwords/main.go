package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chooban/aoc2020/internal/passwords"
)

func main() {
	fmt.Println("Reading", os.Args[1])
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")

	validCount := 0
	for _, v := range lines {
		if len(v) == 0 {
			continue
		}
		values := strings.Split(v, ":")

		fmt.Printf("Validating %s\n", v)
		valid := passwords.ValidateTobogganPassword(values[0], values[1])

		if valid {
			validCount++
		} else {
		}
	}

	fmt.Printf("%d valid passwords", validCount)
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/chooban/aoc2020/internal/io"
	"github.com/chooban/aoc2020/internal/passports"
)

func main() {
	lines := io.ReadFileLines(os.Args[1])

	passportsInput := [][]string{{}}

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			passportsInput[len(passportsInput)-1] = append(passportsInput[len(passportsInput)-1], lines[i])
		} else {
			if i != len(lines)-1 {
				passportsInput = append(passportsInput, []string{})
			}
		}
	}

	for i, passport := range passportsInput {
		updatedPassport := []string{}
		for _, passportLines := range passport {
			fields := strings.Split(passportLines, " ")
			updatedPassport = append(updatedPassport, fields[:]...)
		}
		sort.Strings(updatedPassport)
		passportsInput[i] = updatedPassport
	}

	validCount := 0
	for _, pp := range passportsInput {
		if passports.ValidatePassport(pp) {
			fmt.Printf("%s was valid!\n", pp)
			validCount++
		}
	}
	fmt.Println(validCount)
}

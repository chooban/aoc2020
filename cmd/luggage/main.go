package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/chooban/aoc2020/internal/io"
)

var wholeLineRegex = regexp.MustCompile("^(\\w+ \\w+) bags? contain (.*)\\.$")
var childRegex = regexp.MustCompile("(\\d+) (\\w+ \\w+) bags?")

type BagRule struct {
	Colour string
	Count  int
}

func main() {
	lines := io.ReadFileLines(os.Args[1])

	rules := make(map[string][]BagRule)

	for _, line := range lines {
		if line == "" {
			continue
		}

		found := wholeLineRegex.FindAllStringSubmatch(line, -1)[0]

		if found[2] != "no other bags" {
			childrenInput := strings.Split(found[2], ",")
			children := make([]BagRule, len(childrenInput))

			for idx, child := range childrenInput {
				childParts := childRegex.FindAllStringSubmatch(child, -1)[0]
				count, _ := strconv.Atoi(childParts[1])
				rule := BagRule{
					Colour: string(childParts[2]),
					Count:  count,
				}
				children[idx] = rule
			}

			rules[found[1]] = children
		}

	}

	bagCount := 0
	for c := range rules {
		if containsBag(rules, c, "shiny gold") {
			bagCount++
		}
	}

	fmt.Printf("%d bags can contain a shiny gold\n", bagCount)

	counts := make(map[string]int)

	fmt.Printf("shiny gold bags contains %d bags\n", containsTotalBags(counts, rules, "shiny gold"))
}

func containsBag(bags map[string][]BagRule, rootColour, searchColour string) bool {
	contents := bags[rootColour]

	for _, rule := range contents {
		if rule.Colour == searchColour {
			return true
		}
	}

	for _, rule := range contents {
		if containsBag(bags, rule.Colour, searchColour) {
			return true
		}
	}
	return false
}

func containsTotalBags(counts map[string]int, bags map[string][]BagRule, searchColour string) int {
	count := 0

	contents := bags[searchColour]

	for _, rule := range contents {
		// Add this many bags to the count
		count += rule.Count

		// Now add ruleCount times how ever many it contains
		count += containsTotalBags(counts, bags, rule.Colour) * rule.Count
	}

	counts[searchColour] = count

	return count
}

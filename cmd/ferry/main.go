package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/chooban/aoc2020/internal/io"
)

const (
	availableSeat int = iota
	occupiedSeat
	floor
)

func main() {
	lines := io.ReadFileLines(os.Args[1])

	rows := len(lines)
	cols := len(strings.TrimSpace(lines[0]))

	floorPlan := make([]string, 0, rows*cols)
	for _, row := range lines {
		elements := strings.Split(row, "")
		floorPlan = append(floorPlan, elements...)
	}

	adjacentSeats := make([][]int, len(floorPlan))

	// These won't change
	for i := range adjacentSeats {
		adjacentSeats[i] = adjacentIndices(i, cols, rows)
	}

	iterations := 0
	for {
		changes := map[int]string{}

		for i, val := range floorPlan {
			switch val {
			case "L":
				noOccupied := true
				for _, adj := range adjacentSeats[i] {
					if floorPlan[adj] == "#" {
						noOccupied = false
						break
					}
				}
				if noOccupied == true {
					// Take the seat
					changes[i] = "#"
				}
			case "#":
				occupiedCount := 0
				for _, adj := range adjacentSeats[i] {
					if floorPlan[adj] == "#" {
						occupiedCount++
					}
				}
				if occupiedCount >= 4 {
					changes[i] = "L"
				}
				break
			case ".":
				break
			default:
				fmt.Println("Unknown value:", val)
				panic("Unknown!")
			}
		}
		iterations++

		if len(changes) == 0 {
			// We're done here
			break
		} else {
			// Apply the changes directly
			for k, v := range changes {
				floorPlan[k] = v
			}
		}
	}

	occupied := 0
	for _, val := range floorPlan {
		if val == "#" {
			occupied++
		}
	}

	fmt.Println("After", iterations, "iterations there are", occupied, "occupied seats")
	// printFloorPlan(floorPlan, cols)

}

func printFloorPlan(plan []string, cols int) {
	for i := 0; i < len(plan); i += cols {
		batch := plan[i:min(i+cols, len(plan))]
		fmt.Println(batch)
	}
	fmt.Println()
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func adjacentIndices(i, cols, rows int) []int {
	maxIdx := cols * rows
	if i >= maxIdx {
		return nil
	}
	rawIndices := []int{}
	if i%cols == 0 {
		rawIndices = []int{
			i - cols,
			i - (cols - 1),
			i + 1,
			i + (cols + 1),
			i + cols,
		}

	} else if i%cols == cols-1 {
		rawIndices = []int{
			i - (cols + 1),
			i - cols,
			i - 1,
			i + (cols - 1),
			i + cols,
		}
	} else {
		rawIndices = []int{
			i + (cols + 1),
			i + (cols - 1),
			i + 1,
			i + cols,
			i - (cols + 1),
			i - (cols - 1),
			i - 1,
			i - cols,
		}
	}

	indices := []int{}
	for _, v := range rawIndices {
		if v >= 0 && v < maxIdx {
			indices = append(indices, v)
		}
	}

	sort.Ints(indices)
	return indices
}

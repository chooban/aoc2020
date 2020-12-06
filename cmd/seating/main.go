package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/chooban/aoc2020/internal/io"
)

func main() {
	lines := io.ReadFileLines(os.Args[1])

	seatIDs := []int{}

	for _, bp := range lines {
		if len(bp) == 0 {
			continue
		}
		parts := strings.Split(bp, "")

		row := parts[0:7]
		seat := parts[7:10]

		minRow := 0
		maxRow := 127
		for _, rowPart := range row {
			switch rowPart {
			case "F":
				maxRow = maxRow - ((maxRow - minRow) / 2) - 1
			case "B":
				minRow = minRow + ((maxRow - minRow) / 2) + 1
			}
		}

		if minRow != maxRow {
			panic("Rows didn't coalesce!")
		}

		minSeat := 0
		maxSeat := 7
		for _, seatPart := range seat {
			switch seatPart {
			case "L":
				maxSeat = maxSeat - ((maxSeat - minSeat) / 2) - 1
			case "R":
				minSeat = minSeat + ((maxSeat - minSeat) / 2) + 1
			}
		}

		if minSeat != maxSeat {
			panic("Seats did not coalesce!")
		}

		seatID := minRow*8 + minSeat
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)

	fmt.Print("Max seat ID: ")
	fmt.Println(seatIDs[len(seatIDs)-1])

	for idx := 1; idx < len(seatIDs)-1; idx++ {
		if seatIDs[idx-1] != seatIDs[idx]-1 ||
			seatIDs[idx+1] != seatIDs[idx]+1 {
			fmt.Printf("My seat is %d\n", seatIDs[idx]+1)
			break
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

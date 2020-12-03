package toboggan

func HowManyTrees(slopes []string, moveRight, moveDown int) int {
	row, column, trees := 0, 0, 0

	for ; row < len(slopes); row += moveDown {
		if len(slopes[row]) == 0 {
			continue
		}
		location := string(slopes[row][column])
		if location == "#" {
			trees++
		}
		column = (column + moveRight) % len(slopes[row])
	}
	return trees
}

package expenses

func FindElementsThatSum(expenses []int, target int, size int) []int {
	if size < 2 {
		panic("Bad input")
	}

	possibles := [][]int{}
	for _, val := range expenses {
		if val < target {
			possibles = append(possibles, []int{val})
		}
	}

	for i := 1; i < size; i++ {
		newPossibles := [][]int{}
		for _, possible := range possibles {
			for _, expense := range expenses {
				if contains(possible, expense) {
					continue
				}

				toCheck := append(possible, expense)

				if i == size-1 && sum(toCheck) == target {
					return toCheck
				}

				if i < size-1 && sum(toCheck) < target {
					newPossibles = append(newPossibles, toCheck)
				}
			}
		}

		if len(newPossibles) > 0 {
			possibles = newPossibles
		} else {
			return nil
		}
	}

	return nil
}

func sum(array []int) int {
	sum := 0
	for _, val := range array {
		sum += val
	}
	return sum
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

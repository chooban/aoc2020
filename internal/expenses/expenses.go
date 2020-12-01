package expenses

func FindSumToPair(expenses []int, target int) []int {
	for idxA, valueA := range expenses {
		for idxB, valueB := range expenses {
			if idxA == idxB {
				continue
			}

			if valueA+valueB == target {
				return []int{valueA, valueB}
			}
		}
	}
	return nil
}

func FindSumToTriple(expenses []int, target int) []int {
	for idxA, valueA := range expenses {
		for idxB, valueB := range expenses {
			if idxA == idxB {
				continue
			}

			for idxC, valueC := range expenses {
				if idxC == idxA || idxC == idxB {
					continue
				}
				if valueA+valueB+valueC == target {
					return []int{valueA, valueB, valueC}
				}
			}
		}
	}
	return nil
}

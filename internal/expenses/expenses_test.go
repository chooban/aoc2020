package expenses

import (
	"reflect"
	"testing"
)

func TestSumToTarget(t *testing.T) {
	expected := []int{1, 3}
	total := FindElementsThatSum([]int{1, 3, 4, 5}, 4, 2)

	if !reflect.DeepEqual(expected, total) {
		t.Errorf("Pair was incorrect, got {%d, %d}", total[0], total[1])
	}
}

func TestSumToTargetAgain(t *testing.T) {
	expected := []int{1, 3, 4}
	total := FindElementsThatSum([]int{1, 3, 4, 5}, 8, 3)

	if total == nil {
		t.Errorf("No values found")
	}

	if !reflect.DeepEqual(expected, total) {
		t.Errorf("Pair was incorrect, got {%d, %d, %d}", total[0], total[1], total[2])
	}
}

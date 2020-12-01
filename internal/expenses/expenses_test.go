package expenses

import (
	"reflect"
	"testing"
)

func TestSumTo(t *testing.T) {
	expected := []int{1, 3}
	total := FindSumToPair([]int{1, 3, 4, 5}, 4)

	if !reflect.DeepEqual(expected, total) {
		t.Errorf("Pair was incorrect, got {%d, %d}", total[0], total[1])
	}
}

func TestSumToTriple(t *testing.T) {
	expected := []int{1, 3, 4}
	total := FindSumToTriple([]int{1, 3, 4, 5}, 8)

	if !reflect.DeepEqual(expected, total) {
		t.Errorf("Pair was incorrect, got {%d, %d, %d}", total[0], total[1], total[2])
	}
}

package algorithm

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	assertEqual := func(val []int, exp []int) {
		if len(val) != len(exp) {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
		for i := range val {
			if val[i] != exp[i] {
				t.Errorf("Expected %v, got %v.", exp, val)
			}
		}
	}

	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	MergeSort(array, 0, len(array)-1)
	assertEqual(array, exp)

	array = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	MergeSort(array, 0, len(array)-1)
	assertEqual(array, exp)

	array = []int{3, 2, 4, 7, 1, 9, 8, 6, 0, 5}
	MergeSort(array, 0, len(array)-1)
	assertEqual(array, exp)

	exp = []int{0, 1, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9, 9}
	array = []int{3, 9, 2, 4, 7, 3, 1, 9, 8, 6, 7, 0, 5}
	MergeSort(array, 0, len(array)-1)
	assertEqual(array, exp)
}

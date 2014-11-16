package algorithm

import (
	"testing"
)

func TestFindMaxSubarray(t *testing.T) {
	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	array := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	low, high, sum := FindMaxSubarray(array, 0, len(array)-1)
	assertEqual(low, 7)
	assertEqual(high, 10)
	assertEqual(sum, 43)

	array = []int{1}
	low, high, sum = FindMaxSubarray(array, 0, len(array)-1)
	assertEqual(low, 0)
	assertEqual(high, 0)
	assertEqual(sum, 1)

	array = []int{1, 2, 3}
	low, high, sum = FindMaxSubarray(array, 0, len(array)-1)
	assertEqual(low, 0)
	assertEqual(high, 2)
	assertEqual(sum, 6)

	array = []int{-1, -2, -3}
	low, high, sum = FindMaxSubarray(array, 0, len(array)-1)
	assertEqual(low, 0)
	assertEqual(high, 0)
	assertEqual(sum, -1)
}

package algorithm

import (
	"testing"
)

func TestCutRod(t *testing.T) {
	assertEqual := func(val int, exp int) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	maxRevenue := map[int]int{
		1:  1,
		2:  5,
		3:  8,
		4:  10,
		5:  13,
		6:  17,
		7:  18,
		8:  22,
		9:  25,
		10: 30,
	}
	for length, exp := range maxRevenue {
		revenue, _ := CutRod(length, price)
		assertEqual(revenue, exp)
	}
}

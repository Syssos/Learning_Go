package main

import "testing"

func TestSum(t *testing.T) {
	cases := []struct{
		x int
		y int
		sum int
	}{
		{4, 1, 5},
		{5, 2, 7},
		{0, 2, 2},
	}

	for _, i := range cases {
		res := Sum(i.x, i.y)
		if res != i.sum {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", i.x, i.y, res, i.sum)
		}
	}
}
package main

import (
	"testing"
)

var Result int

func TestSum(t *testing.T) {
	cases := []struct{
		x int
		y int
		sum int
	}{
		{4, 1, 5},
		{5, 2, 7},
		// {9223372036854775807, 1, 9223372036854775808},
	}
	for _, i := range cases {
		res := Sum(i.x, i.y)
		if res != i.sum {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", i.x, i.y, res, i.sum)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
        r = Sum(4, 5)
    }

    Result = r
}
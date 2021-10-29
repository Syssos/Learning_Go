package main

import (
	"testing"
)

func TestDif(t *testing.T) {
	cases := []struct{
		x int
		y int
		sum int
	}{
		{4, 1, 3},
		{5, 2, 3},
		// {9223372036854775807, 1, 9223372036854775808},
	}
	for _, i := range cases {
		res := Dif(i.x, i.y)
		if res != i.sum {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", i.x, i.y, res, i.sum)
		}
	}
}

func BenchmarkDif(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
        r = Dif(5, 4)
    }

    Result = r
}
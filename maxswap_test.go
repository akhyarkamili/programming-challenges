package main

import "testing"

func Test_MaxSwapReturnsZero(t *testing.T) {
	tt := []struct {
		name string
		n    int
		want int
	}{
		{
			"n is zero",
			0,
			0,
		},
		{
			"n is single digit",
			1,
			1,
		},
		{
			"swap is needed",
			12,
			21,
		},
		{
			"swap with two bigger number",
			1299,
			9291,
		},
		{
			"swap not the leftmost",
			9299,
			9992,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := getMaximumSwap(tc.n)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

package leet_743

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNetworkDelayTime(t *testing.T) {
	var testCase = []struct {
		times  [][]int
		n, k   int
		expect int
	}{
		{
			times: [][]int{
				[]int{1, 2, 1},
			},
			n:      2,
			k:      2,
			expect: -1,
		},
		{
			times: [][]int{
				[]int{2, 1, 1},
				[]int{2, 3, 1},
				[]int{3, 4, 1},
			},
			n:      4,
			k:      2,
			expect: 2,
		},
		{
			times: [][]int{
				[]int{1, 2, 2},
				[]int{1, 3, 5},
				[]int{2, 3, 2},
				[]int{2, 4, 8},
				[]int{3, 4, 3},
			},
			n:      4,
			k:      1,
			expect: 7,
		},
		{
			times: [][]int{
				[]int{3, 5, 78},
				[]int{2, 1, 1},
				[]int{1, 3, 0},
				[]int{4, 3, 59},
				[]int{5, 3, 85},
				[]int{5, 2, 22},
				[]int{2, 4, 23},
				[]int{1, 4, 43},
				[]int{4, 5, 75},
				[]int{5, 1, 15},
				[]int{1, 5, 91},
				[]int{4, 1, 16},
				[]int{3, 2, 98},
				[]int{3, 4, 22},
				[]int{5, 4, 31},
				[]int{1, 2, 0},
				[]int{2, 5, 4},
				[]int{4, 2, 51},
				[]int{3, 1, 36},
				[]int{2, 3, 59},
			},
			n:      5,
			k:      5,
			expect: 31,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := networkDelayTime(tc.times, tc.n, tc.k)
		ass.Equal(tc.expect, actual, "times: %v n: %d k: %d", tc.times, tc.n, tc.k)
	}
}

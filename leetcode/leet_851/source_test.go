package leet_851

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoudAndRich(t *testing.T) {
	var testCase = []struct {
		richer [][]int
		quiet  []int
		expect []int
	}{
		{
			richer: [][]int{
				[]int{1, 0},
				[]int{2, 1},
				[]int{3, 1},
				[]int{3, 7},
				[]int{4, 3},
				[]int{5, 3},
				[]int{6, 3},
			},
			quiet:  []int{3, 2, 5, 4, 6, 1, 7, 0},
			expect: []int{5, 5, 2, 5, 4, 5, 6, 7},
		},
	}
	ass := require.New(t)
	for idx, tc := range testCase {
		actual := loudAndRich(tc.richer, tc.quiet)
		ass.Equal(tc.expect, actual, "#%d failed", idx)
	}
}

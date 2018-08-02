package leet_769

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxChunksToSorted(t *testing.T) {
	var testCase = []struct {
		arr    []int
		expect int
	}{
		{
			arr:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expect: 10,
		},
		{
			arr:    []int{0},
			expect: 1,
		},
		{
			arr:    []int{1, 0, 9, 8, 7, 6, 5, 4, 3, 2},
			expect: 2,
		},
		{
			arr:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expect: 1,
		},
		{
			arr:    []int{1, 2, 0, 4, 5, 3, 6},
			expect: 3,
		},
		{
			arr:    []int{1, 0, 2, 3, 4},
			expect: 4,
		},
		{
			arr:    []int{4, 3, 2, 1, 0},
			expect: 1,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := maxChunksToSorted(tc.arr)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

package leet_84

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestRectangleArea(t *testing.T) {
	var testCase = []struct {
		heights []int
		expect  int
	}{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			expect:  10,
		},
		{
			heights: []int{2, 3, 3, 3, 2},
			expect:  10,
		},
		{
			heights: []int{2, 3, 3, 3, 1},
			expect:  9,
		},
		{
			heights: []int{6},
			expect:  6,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := largestRectangleArea(tc.heights)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

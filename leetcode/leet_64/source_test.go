package leet_64

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinPathSum(t *testing.T) {
	var testCase = []struct {
		grid   [][]int
		expect int
	}{
		{
			grid: [][]int{
				[]int{0},
			},
			expect: 0,
		},
		{
			grid: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			expect: 7,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := minPathSum(tc.grid)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

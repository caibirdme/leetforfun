package leet_892

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSurfaceArea(t *testing.T) {
	var testCase = []struct {
		grid   [][]int
		expect int
	}{
		{
			grid: [][]int{
				[]int{2},
			},
			expect: 10,
		},
		{
			grid: [][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			expect: 34,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := surfaceArea(tc.grid)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

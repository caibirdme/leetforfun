package leet_887

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectionArea(t *testing.T) {
	var testCase = []struct {
		grid   [][]int
		expect int
	}{
		{
			grid: [][]int{
				[]int{2},
			},
			expect: 5,
		},
		{
			grid: [][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			expect: 17,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := projectionArea(tc.grid)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

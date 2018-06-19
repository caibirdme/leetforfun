package leet_63

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	var testData = []struct {
		grid   [][]int
		expect int
	}{
		{
			grid: [][]int{
				[]int{0, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			expect: 2,
		},
	}
	ass := require.New(t)
	for _, tc := range testData {
		actual := uniquePathsWithObstacles(tc.grid)
		ass.Equal(tc.expect, actual)
	}
}

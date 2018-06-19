package leet_547

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindCircleNum(t *testing.T) {
	var testCase = []struct {
		M      [][]int
		expect int
	}{
		{
			M: [][]int{
				[]int{1, 1, 0},
				[]int{1, 1, 0},
				[]int{0, 0, 1},
			},
			expect: 2,
		},
		{
			M: [][]int{
				[]int{1, 1, 0},
				[]int{1, 1, 1},
				[]int{0, 1, 1},
			},
			expect: 1,
		},
		{
			M: [][]int{
				[]int{1, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 1},
			},
			expect: 3,
		},
	}
	ass := require.New(t)
	for idx, tc := range testCase {
		actual := findCircleNum(tc.M)
		ass.Equal(tc.expect, actual, "idx#%d failed", idx)
	}
}

package leet_886

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReachableNodes(t *testing.T) {
	var testCase = []struct {
		edges  [][]int
		M, N   int
		expect int
	}{
		{
			edges: [][]int{
				[]int{0, 2, 0},
				[]int{1, 3, 1},
				[]int{0, 1, 0},
				[]int{1, 4, 0},
				[]int{0, 4, 0},
				[]int{2, 4, 4},
				[]int{2, 3, 6},
				[]int{0, 3, 8},
				[]int{3, 4, 1},
				[]int{1, 2, 4},
			},
			M:      4,
			N:      5,
			expect: 24,
		},
		{
			edges: [][]int{
				[]int{2, 4, 2},
				[]int{3, 4, 5},
				[]int{2, 3, 1},
				[]int{0, 2, 1},
				[]int{0, 3, 5},
			},
			M:      14,
			N:      5,
			expect: 18,
		},
		{
			edges: [][]int{
				[]int{0, 1, 4},
				[]int{1, 2, 6},
				[]int{0, 2, 8},
				[]int{1, 3, 1},
			},
			M:      10,
			N:      4,
			expect: 23,
		},
		{
			edges: [][]int{
				[]int{0, 1, 10},
				[]int{0, 2, 1},
				[]int{1, 2, 2},
			},
			M:      6,
			N:      3,
			expect: 13,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := reachableNodes(tc.edges, tc.M, tc.N)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

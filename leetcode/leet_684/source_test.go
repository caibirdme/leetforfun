package leet_684

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindRedundantConnection(t *testing.T) {
	var testCase = []struct {
		edges  [][]int
		expect []int
	}{
		{
			edges: [][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{2, 3},
			},
			expect: []int{2, 3},
		},
		{
			edges: [][]int{
				[]int{1, 2},
				[]int{2, 3},
				[]int{3, 4},
				[]int{1, 4},
				[]int{1, 5},
			},
			expect: []int{1, 4},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := findRedundantConnection(tc.edges)
		should.Equal(tc.expect, actual, "case#%d failed", idx)
	}
}

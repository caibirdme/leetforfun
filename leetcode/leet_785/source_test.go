package leet_785

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBipartite(t *testing.T) {
	var testCase = []struct {
		graph  [][]int
		expect bool
	}{
		{
			graph: [][]int{
				[]int{}, []int{2, 4, 6}, []int{1, 4, 8, 9}, []int{7, 8}, []int{1, 2, 8, 9}, []int{6, 9}, []int{1, 5, 7, 8, 9}, []int{3, 6, 9}, []int{2, 3, 4, 6, 9}, []int{2, 4, 5, 6, 7, 8},
			},
			expect: false,
		},
		{
			graph: [][]int{
				[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2},
			},
			expect: true,
		},
		{
			graph: [][]int{
				[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2},
			},
			expect: false,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := isBipartite(tc.graph)
		should.Equal(tc.expect, actual, "case#%d failed", idx)
	}
}

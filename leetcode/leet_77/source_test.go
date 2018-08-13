package leet_77

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombine(t *testing.T) {
	var testCase = []struct {
		n, k   int
		expect [][]int
	}{
		{
			n: 1,
			k: 1,
			expect: [][]int{
				[]int{1},
			},
		},
		{
			n: 3,
			k: 3,
			expect: [][]int{
				[]int{1, 2, 3},
			},
		},
		{
			n: 4,
			k: 3,
			expect: [][]int{
				[]int{1, 2, 3},
				[]int{1, 2, 4},
				[]int{1, 3, 4},
				[]int{2, 3, 4},
			},
		},
		{
			n: 4,
			k: 2,
			expect: [][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 4},
			},
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := combine(tc.n, tc.k)
		should.Equal(tc.expect, actual, "n:%d k:%d", tc.n, tc.k)
	}
}

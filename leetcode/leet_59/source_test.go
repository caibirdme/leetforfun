package leet_59

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMatrix(t *testing.T) {
	var testCase = []struct {
		n      int
		expect [][]int
	}{
		{
			n: 3,
			expect: [][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			n: 4,
			expect: [][]int{
				[]int{1, 2, 3, 4},
				[]int{12, 13, 14, 5},
				[]int{11, 16, 15, 6},
				[]int{10, 9, 8, 7},
			},
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := generateMatrix(tc.n)
		should.Equal(tc.expect, actual, "n=%d", tc.n)
	}
}

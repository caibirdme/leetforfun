package leet_873

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLenLongestFibSubseq(t *testing.T) {
	var testCase = []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			expect: 5,
		},
		{
			A:      []int{1, 3, 7, 11, 12, 14, 18},
			expect: 3,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := lenLongestFibSubseq(tc.A)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

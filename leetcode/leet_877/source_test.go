package leet_877

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoneGame(t *testing.T) {
	var testCase = []struct {
		piles  []int
		expect bool
	}{
		{
			piles:  []int{2, 3},
			expect: true,
		},
		{
			piles:  []int{2, 3, 10, 1},
			expect: true,
		},
		{
			piles:  []int{5, 3, 4, 5},
			expect: true,
		},
		{
			piles:  []int{2, 3, 4, 5},
			expect: true,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := stoneGame(tc.piles)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

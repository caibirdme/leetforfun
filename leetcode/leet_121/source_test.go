package leet_121

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	var testCase = []struct {
		prices []int
		expect int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		}, {
			prices: []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := maxProfit(tc.prices)
		should.Equal(tc.expect, actual, "case# %d fail", idx)
	}
}

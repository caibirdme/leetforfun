package leet_122

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
			prices: []int{1},
			expect: 0,
		},
		{
			prices: []int{1, 2},
			expect: 1,
		},
		{
			prices: []int{1, 3, 1, 3, 5},
			expect: 6,
		},
		{
			prices: []int{1, 3, 1, 3, 2},
			expect: 4,
		},
		{
			prices: []int{1, 3, 1, 3, 2, 3},
			expect: 5,
		},
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			expect: 7,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			expect: 0,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			expect: 4,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := maxProfit(tc.prices)
		should.Equal(tc.expect, actual, "case# %d fail", idx)
	}
}

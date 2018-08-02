package leet_188

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	var testCase = []struct {
		prices []int
		k      int
		expect int
	}{
		{
			prices: []int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9},
			k:      2,
			expect: 16,
		},
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			k:      1,
			expect: 5,
		},
		{
			prices: []int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6},
			k:      2,
			expect: 11,
		},
		{
			prices: []int{5, 5, 4, 9, 3, 8, 5, 5, 1, 6, 8, 3, 4},
			k:      2,
			expect: 12,
		},

		{
			prices: []int{2, 4, 1},
			k:      2,
			expect: 2,
		},
		{
			prices: []int{3, 2, 6, 5, 0, 3},
			k:      2,
			expect: 7,
		},
		{
			prices: []int{1, 2, 1, 2, 1, 2},
			k:      3,
			expect: 3,
		},
		{
			prices: []int{1, 0, 2, 3, 2, 2, 4},
			k:      3,
			expect: 5,
		},
		{
			prices: []int{1, 2, 3},
			k:      0,
			expect: 0,
		},
		{
			prices: []int{8, 7, 6, 5, 4, 3, 2, 1},
			k:      5,
			expect: 0,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := maxProfit(tc.k, tc.prices)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

package leet_123

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
			prices: []int{1, 4, 2, 7},
			expect: 8,
		},
		{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			expect: 6,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			expect: 4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := maxProfit(tc.prices)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

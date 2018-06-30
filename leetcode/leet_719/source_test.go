package leet_719

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallestDistancePair(t *testing.T) {
	var testCase = []struct {
		nums   []int
		k      int
		expect int
	}{
		{
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6},
			k:      8,
			expect: 1,
		},
		{
			nums:   []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			k:      1,
			expect: 0,
		},
		{
			nums:   []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			k:      5,
			expect: 2,
		},
		{
			nums:   []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			k:      4,
			expect: 1,
		},
		{
			nums:   []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			k:      3,
			expect: 1,
		},
		{
			nums:   []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			k:      2,
			expect: 1,
		},
		{
			nums:   []int{1, 5},
			k:      1,
			expect: 4,
		},
		{
			nums:   []int{1, 3, 1},
			k:      1,
			expect: 0,
		},
		{
			nums:   []int{1, 3, 7, 20, 6},
			k:      1,
			expect: 1,
		},
		{
			nums:   []int{1, 3, 7, 20, 6},
			k:      2,
			expect: 2,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := smallestDistancePair(tc.nums, tc.k)
		should.Equal(tc.expect, actual, "case#%d failed", idx)
	}
}

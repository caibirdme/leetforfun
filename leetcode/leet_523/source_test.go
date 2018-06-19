package leet_523

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckSubarraySum(t *testing.T) {
	var testCase = []struct {
		nums   []int
		k      int
		expect bool
	}{
		{
			nums:   []int{23, 2, 4, 6, 7},
			k:      6,
			expect: true,
		},
		{
			nums:   []int{23, 2, 4, 6, 7},
			k:      0,
			expect: false,
		},
		{
			nums:   []int{23, 2, 4, 6, 7},
			k:      1,
			expect: true,
		},
		{
			nums:   []int{0, 0},
			k:      1,
			expect: true,
		},
		{
			nums:   []int{0, 0},
			k:      0,
			expect: true,
		},
		{
			nums:   []int{0},
			k:      0,
			expect: false,
		},
		{
			nums:   []int{1},
			k:      1,
			expect: false,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		ass.Equal(tc.expect, checkSubarraySum(tc.nums, tc.k), "k: %d, nums: %v", tc.k, tc.nums)
	}
}

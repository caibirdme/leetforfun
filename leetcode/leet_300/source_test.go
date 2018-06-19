package leet_300

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLIS(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{1, 1},
			expect: 1,
		},
		{
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			expect: 4,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			expect: 8,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 7, 7, 8},
			expect: 7,
		},
		{
			nums:   []int{1, 2, 3, 4, 7, 7, 7, 8},
			expect: 6,
		},
		{
			nums:   []int{1, 2, 3, 4, 0, 1, 2, 3, 4},
			expect: 5,
		},
		{
			nums:   []int{1},
			expect: 1,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := lengthOfLIS(tc.nums)
		ass.Equal(tc.expect, actual, "nums: %v", tc.nums)
	}
}

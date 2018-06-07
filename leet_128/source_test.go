package leet_128

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestConsecutive(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{-4, -1, -5, 1, -6, 0, 2, -3, -2, 3},
			expect: 10,
		},
		{
			nums:   []int{-4, 0, -1, -2, 1, -3},
			expect: 6,
		},
		{
			nums:   []int{-4, -1, 1, 0, 2, -3, -2},
			expect: 7,
		},
		{
			nums:   []int{-4, -1, 4, -5, 1, -6, 9, -6, 0, 2, 2, 7, 0, 9, -3, 8, 9, -2, -6, 5, 0, 3, 4, -2},
			expect: 12,
		},
		{
			nums:   []int{-2, 0, -1},
			expect: 3,
		},
		{
			nums:   []int{0, 3, 2, 5, 4, 6, 0, 1},
			expect: 7,
		},
		{
			nums:   []int{0, 3, 2, 0, 1},
			expect: 4,
		},
		{
			nums:   []int{0},
			expect: 1,
		},
		{
			nums:   []int{},
			expect: 0,
		},
		{
			nums:   []int{100, 98, 4},
			expect: 1,
		},
		{
			nums:   []int{100, 98, 4, 23, 3, 1, 2},
			expect: 4,
		},
		{
			nums:   []int{5, 100, 98, 4, 23, 3, 1, 2},
			expect: 5,
		},
		{
			nums:   []int{5, 100, 98, 4, 23, 3, 1, 2, 99, 98, 97, 96, 96, 95, 94},
			expect: 7,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := longestConsecutive(tc.nums)
		ass.Equal(tc.expect, actual, "nums: %v", tc.nums)
	}
}

package leet_781

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumRabbits(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{100},
			expect: 101,
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
			nums:   []int{2, 2, 2, 2, 2, 3, 4, 4, 4, 4, 4, 4, 1},
			expect: 22,
		},
		{
			nums:   []int{1, 1, 2},
			expect: 5,
		},
		{
			nums:   []int{10, 10, 10},
			expect: 11,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		ass.Equal(tc.expect, numRabbits(tc.nums), "nums: %v", tc.nums)
	}
}

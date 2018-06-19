package leet_560

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubarraySum(t *testing.T) {
	var testCase = []struct {
		nums   []int
		k      int
		expect int
	}{
		{
			nums:   []int{1, 1, 1},
			k:      2,
			expect: 2,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := subarraySum(tc.nums, tc.k)
		ass.Equal(tc.expect, actual, "nums: %v k:%d", tc.nums, tc.k)
	}
}

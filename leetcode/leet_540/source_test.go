package leet_540

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNonDuplicate(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{1, 1, 2, 3, 3, 4, 4},
			expect: 2,
		},
		{
			nums:   []int{1, 1, 2, 2, 3, 3, 5, 4, 4},
			expect: 5,
		},
		{
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 4, 4},
			expect: 4,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		should.Equal(tc.expect, singleNonDuplicate(tc.nums), "nums: %v", tc.nums)
	}
}

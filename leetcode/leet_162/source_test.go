package leet_162

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindPeakElement(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{5},
			expect: 0,
		},
		{
			nums:   []int{1, 2, 3, 1},
			expect: 2,
		},
		{
			nums:   []int{1, 2, 1, 3, 5, 6, 4},
			expect: 1,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := findPeakElement(tc.nums)
		should.Equal(tc.expect, actual, "%v", tc.nums)
	}
}

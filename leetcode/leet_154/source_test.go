package leet_154

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMin(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{3, 2, 1},
			expect: 1,
		},
		{
			nums:   []int{1, 2, 3},
			expect: 1,
		},
		{
			nums:   []int{2, 2, 2, 0, 1},
			expect: 0,
		},
		{
			nums:   []int{4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 2, 4, 4},
			expect: 1,
		},
		{
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
			expect: 1,
		},
		{
			nums:   []int{4, 4, 1, 1, 1, 1, 1, 1, 4},
			expect: 1,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := findMin(tc.nums)
		should.Equal(tc.expect, actual, "%v", tc.nums)
	}
}

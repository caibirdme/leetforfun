package leet_137

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumber(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{-2, -2, 3, -2},
			expect: 3,
		},
		{
			nums:   []int{2, 2, 3, 2},
			expect: 3,
		},
		{
			nums:   []int{0, 1, 0, 1, 0, 1, 99},
			expect: 99,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := singleNumber(tc.nums)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

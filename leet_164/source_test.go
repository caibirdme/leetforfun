package leet_164

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumGap(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{3, 6, 9, 1},
			expect: 3,
		},
		{
			nums:   []int{1, 3, 8, 2, 20},
			expect: 12,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := maximumGap(tc.nums)
		ass.Equal(tc.expect, actual)
	}
}

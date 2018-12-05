package leet_189

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	var testCase = []struct {
		nums   []int
		k      int
		expect []int
	}{
		{
			nums:   []int{-1, -100, 3, 99},
			k:      2,
			expect: []int{3, 99, -1, -100},
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			expect: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      6,
			expect: []int{2, 3, 4, 5, 6, 7, 1},
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      14,
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      17,
			expect: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		rotate(tc.nums, tc.k)
		should.Equal(tc.expect, tc.nums, "case #%d fail", idx)
	}
}

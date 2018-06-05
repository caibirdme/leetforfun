package leet_209

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinSubArrayLen(t *testing.T) {
	var testCase = []struct {
		s      int
		nums   []int
		expect int
	}{
		{
			s:      11,
			nums:   []int{2, 1, 1, 1, 1, 1, 1, 1, 1, 10000},
			expect: 1,
		},
		{
			s:      10,
			nums:   []int{2, 1, 1, 1, 1, 1, 1, 1, 1, 10000},
			expect: 1,
		},
		{
			s:      213,
			nums:   []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			expect: 8,
		},
		{
			s:      7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			expect: 2,
		},
		{
			s:      4,
			nums:   []int{2, 1, 1, 2},
			expect: 3,
		},
		{
			s:      4,
			nums:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expect: 0,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := minSubArrayLen(tc.s, tc.nums)
		ass.Equal(tc.expect, actual, "s: %d nums:%v", tc.s, tc.nums)
	}
}

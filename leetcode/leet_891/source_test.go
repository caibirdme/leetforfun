package leet_891

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumSubseqWidths(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{5, 69, 89, 92, 31, 16, 25, 45, 63, 40, 16, 56, 24, 40, 75, 82, 40, 12, 50, 62, 92, 44, 67, 38, 92, 22, 91, 24, 26, 21, 100, 42, 23, 56, 64, 43, 95, 76, 84, 79, 89, 4, 16, 94, 16, 77, 92, 9, 30, 13},
			expect: 857876214,
		},
		{
			nums:   []int{96, 87, 191, 197, 40, 101, 108, 35, 169, 50, 168, 182, 95, 80, 144, 43, 18, 60, 174, 13, 77, 173, 38, 46, 80, 117, 13, 19, 11, 6, 13, 118, 39, 80, 171, 36, 86, 156, 165, 190, 53, 49, 160, 192, 57, 42, 97, 35, 124, 200, 84, 70, 145, 180, 54, 141, 159, 42, 66, 66, 25, 95, 24, 136, 140, 159, 71, 131, 5, 140, 115, 76, 151, 137, 63, 47, 69, 164, 60, 172, 153, 183, 6, 70, 40, 168, 133, 45, 116, 188, 20, 52, 70, 156, 44, 27, 124, 59, 42, 172},
			expect: 136988321,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			expect: 72,
		},
		{
			nums:   []int{1, 1, 4},
			expect: 9,
		},
		{
			nums:   []int{3, 7, 2, 3},
			expect: 35,
		},
		{
			nums:   []int{2, 1, 3},
			expect: 6,
		},

		{
			nums:   []int{2, 1, 4, 3},
			expect: 23,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := sumSubseqWidths(tc.nums)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

package leet_765

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinSwapsCouples(t *testing.T) {
	var testCase = []struct {
		row    []int
		expect int
	}{
		{
			row:    []int{1, 4, 0, 5, 8, 7, 6, 3, 2, 9},
			expect: 3,
		},
		{
			row:    []int{0, 2, 1, 3},
			expect: 1,
		},
		{
			row:    []int{3, 2, 0, 1},
			expect: 0,
		},
		{
			row:    []int{0, 2, 1, 4, 3, 5},
			expect: 2,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := minSwapsCouples(tc.row)
		should.Equal(tc.expect, actual, "row: %v", tc.row)
	}
}

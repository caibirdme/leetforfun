package leet_643

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMaxAverage(t *testing.T) {
	var testCase = []struct {
		nums   []int
		k      int
		expect float64
	}{
		{
			nums:   []int{1, 12, -5, -6, 50, 3},
			k:      4,
			expect: 12.75,
		},
		{
			nums:   []int{1, 2, 3, 4},
			k:      4,
			expect: 2.5,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := findMaxAverage(tc.nums, tc.k)
		should.Equal(tc.expect, actual, "case#%d fail", idx)
	}
}

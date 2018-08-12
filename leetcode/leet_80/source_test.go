package leet_80

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect []int
		length int
	}{
		{
			nums:   []int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
			expect: []int{0, 0, 1, 1, 2, 2},
			length: 6,
		},
		{
			nums:   []int{0, 0, 0},
			expect: []int{0, 0},
			length: 2,
		},
		{
			nums:   []int{1, 1, 2, 2, 2, 3},
			expect: []int{1, 1, 2, 2, 3},
			length: 5,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expect: []int{0, 0, 1, 1, 2, 3, 3},
			length: 7,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		should.Equal(tc.length, removeDuplicates(tc.nums), "case #%d fail", idx)
		should.True(cmp(tc.expect, tc.nums, tc.length), "case #%d fail", idx)
	}
}

func cmp(arr1, arr2 []int, length int) bool {
	if len(arr1) < length || len(arr2) < length {
		return false
	}
	for i := 0; i < length; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

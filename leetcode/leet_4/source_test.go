package leet_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	nums1, nums2 []int
	ans          float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	var data = []testCase{
		testCase{
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		testCase{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		testCase{
			[]int{1, 3, 7, 8, 10, 20},
			[]int{2, 4, 5, 6, 9, 100, 101, 102, 103, 104},
			float64(17) / 2,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.ans, findMedianSortedArrays(tc.nums1, tc.nums2))
	}
}

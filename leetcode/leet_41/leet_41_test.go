package leet_41

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	var testData = []struct {
		in  []int
		out int
	}{
		{
			in:  []int{0, 1, 2},
			out: 3,
		},
		{
			in:  []int{3, 4, -1, 1},
			out: 2,
		},
		{
			in:  []int{-1, -2, -3, -4},
			out: 1,
		},
		{
			in:  []int{-1, -2, -3, -4, 5},
			out: 1,
		},
		{
			in:  []int{-1, -2, -3, -4, 5, 4, 3, 1},
			out: 2,
		},
		{
			in:  []int{-1, -2, -3, -4, 0, 5, 4, 3, 1, 800},
			out: 2,
		},
		{
			in:  []int{},
			out: 1,
		},
		{
			in:  []int{0},
			out: 1,
		},
		{
			in:  []int{1},
			out: 2,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		tar := make([]int, len(tc.in))
		copy(tar, tc.in)
		t := firstMissingPositive(tar)
		ass.Equal(tc.out, t, "input=%v", tc.in)
	}
}

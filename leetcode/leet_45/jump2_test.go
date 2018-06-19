package leet_45

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump2(t *testing.T) {
	var testData = []struct {
		in  []int
		out int
	}{
		{
			in:  []int{3, 0, 0, 1, 1},
			out: 2,
		},
		{
			in:  []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5},
			out: 13,
		},

		{
			in:  []int{1, 1, 1, 6, 1, 100, 1, 1, 1, 1, 1, 1, 1},
			out: 5,
		},
		{
			in:  []int{2, 3, 1, 1, 4},
			out: 2,
		},
		{
			in:  []int{4, 3, 1, 1, 4},
			out: 1,
		},
		{
			in:  []int{1, 1, 1, 1, 1},
			out: 4,
		},
		{
			in:  []int{10000, 1},
			out: 1,
		},
		{
			in:  []int{10000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			out: 1,
		},
		{
			in:  []int{1, 100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			out: 2,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, jump(tc.in), "in: %+v", tc.in)
	}
}

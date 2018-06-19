package leet_42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	var testData = []struct {
		in  []int
		out int
	}{
		{
			in:  []int{3, 0, 3},
			out: 3,
		},
		{
			in:  []int{3, 0, 0, 3},
			out: 6,
		},
		{
			in:  []int{3, 0, 0, 3, 4, 0, 1, 2},
			out: 9,
		},
		{
			in:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			out: 6,
		},
		{
			in:  []int{},
			out: 0,
		},
		{
			in:  []int{1, 2},
			out: 0,
		},
		{
			in:  []int{5, 4, 1, 2},
			out: 1,
		},
		{
			in:  []int{5, 2, 1, 2, 1, 5},
			out: 14,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		area := trap(tc.in)
		ass.Equal(tc.out, area)
	}
}

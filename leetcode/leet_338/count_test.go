package leet_338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	var testData = []struct {
		in  int
		out []int
	}{
		{
			in:  5,
			out: []int{0, 1, 1, 2, 1, 2},
		},
		{
			in:  0,
			out: []int{0},
		},
		{
			in:  1,
			out: []int{0, 1},
		},
		{
			in:  2,
			out: []int{0, 1, 1},
		},
		{
			in:  8,
			out: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, countBits(tc.in))
	}
}

package leet_53

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	var testData = []struct {
		in  []int
		out int
	}{
		{
			in:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			out: 6,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, maxSubArray(tc.in), "in=%+v", tc.in)
	}
}

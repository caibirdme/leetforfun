package leet_54

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	var testData = []struct {
		in  [][]int
		out []int
	}{
		{
			in: [][]int{
				[]int{2, 5},
				[]int{8, 4},
				[]int{0, -1},
			},
			out: []int{2, 5, 4, -1, 0, 8},
		},
		{
			in: [][]int{
				[]int{2, 3},
			},
			out: []int{2, 3},
		},
		{
			in: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			out: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, spiralOrder(tc.in))
	}
}

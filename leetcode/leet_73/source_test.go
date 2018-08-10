package leet_73

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetZeroes(t *testing.T) {
	var testCase = []struct {
		in  [][]int
		out [][]int
	}{
		{
			in: [][]int{
				[]int{1, 1, 1},
				[]int{1, 0, 1},
				[]int{1, 1, 1},
			},
			out: [][]int{
				[]int{1, 0, 1},
				[]int{0, 0, 0},
				[]int{1, 0, 1},
			},
		},
		{
			in: [][]int{
				[]int{0, 1, 2, 0},
				[]int{3, 4, 5, 2},
				[]int{1, 3, 1, 5},
			},
			out: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 4, 5, 0},
				[]int{0, 3, 1, 0},
			},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		setZeroes(tc.in)
		should.Equal(tc.out, tc.in, "case #%d fail", idx)
	}
}

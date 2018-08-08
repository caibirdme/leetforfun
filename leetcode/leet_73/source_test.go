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
	}
	should := require.New(t)
	for idx, tc := range testCase {
		setZeroes(tc.in)
		should.Equal(tc.out, tc.in, "case #%d fail", idx)
	}
}

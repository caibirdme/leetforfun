package leet_85

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalRectangle(t *testing.T) {
	var testData = []struct {
		in  [][]byte
		out int
	}{
		{
			in: [][]byte{
				[]byte{1, 0, 1, 0, 0},
				[]byte{1, 0, 1, 1, 1},
				[]byte{1, 1, 1, 1, 1},
				[]byte{1, 0, 0, 1, 0},
			},
			out: 6,
		},
		{
			in: [][]byte{
				[]byte{1, 0, 1, 0, 0},
				[]byte{1, 1, 1, 1, 1},
				[]byte{1, 1, 1, 1, 1},
				[]byte{1, 0, 0, 1, 0},
			},
			out: 10,
		},
		{
			in:  nil,
			out: 0,
		},
		{
			in: [][]byte{
				[]byte{1},
			},
			out: 1,
		},
		{
			in: [][]byte{
				[]byte{1, 0, 0, 1, 1},
			},
			out: 2,
		},
		{
			in: [][]byte{
				[]byte{1, 0, 0, 1, 1, 1, 0, 1, 1},
			},
			out: 3,
		},
		{
			in: [][]byte{
				[]byte{1, 0, 0, 1, 1, 1, 0, 1, 1},
			},
			out: 3,
		},
		{
			in: [][]byte{
				[]byte{1, 0, 0, 1, 1, 1, 0, 1, 1},
				[]byte{1, 0, 0, 1, 1, 1, 0, 1, 1},
				[]byte{1, 0, 0, 1, 1, 0, 0, 1, 1},
				[]byte{1, 0, 0, 1, 1, 1, 0, 1, 1},
			},
			out: 8,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		n := len(tc.in)
		if n > 0 {
			m := len(tc.in[0])
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					tc.in[i][j] += '0'
				}
			}
		}
		ass.Equal(tc.out, maximalRectangle(tc.in), "input: %v\n", tc.in)
	}
}

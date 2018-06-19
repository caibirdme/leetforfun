package leet_11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	height []int
	ans    int
}

func TestMaxArea(t *testing.T) {
	var data = []Case{
		Case{
			[]int{1, 2},
			1,
		},
		Case{
			[]int{1, 2, 1},
			2,
		},
		Case{
			[]int{1, 2, 3, 4, 5, 6},
			9,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.ans, maxArea(tc.height))
	}
}

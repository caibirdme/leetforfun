package leet_26

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  []int
	out int
}

func TestRemoveDuplicates(t *testing.T) {
	var data = []Case{
		Case{
			[]int{1, 1, 2},
			2,
		},
		Case{
			[]int{1, 1, 1, 2, 2, 2, 3, 4, 4, 5},
			5,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, removeDuplicates(tc.in), "in: %v", tc.in)
	}
}

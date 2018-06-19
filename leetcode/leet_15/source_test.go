package leet_15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	arr []int
	ans [][]int
}

type ClearCase struct {
	arr []int
	ans []int
}

func TestClear(t *testing.T) {
	var data = []ClearCase{
		ClearCase{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0},
		},
		ClearCase{
			[]int{0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 5, 5, 5, 5, 6},
			[]int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 6},
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.ans, clear(tc.arr))
	}
}

func TestThreeSum(t *testing.T) {
	var data = []Case{
		Case{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			},
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		actual := threeSum(tc.arr)
		ass.Equal(tc.ans, actual)
	}
}

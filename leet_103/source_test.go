package leet_103

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const null = -999999

func TestZigzagLevelOrder(t *testing.T) {
	var testData = []struct {
		in  []int
		out [][]int
	}{
		{
			in: []int{3, 9, 20, null, null, 15, 7},
			out: [][]int{
				[]int{3},
				[]int{20, 9},
				[]int{15, 7},
			},
		},
		{
			in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			out: [][]int{
				[]int{1},
				[]int{3, 2},
				[]int{4, 5, 6, 7},
				[]int{10, 9, 8},
			},
		},
		{
			in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			out: [][]int{
				[]int{1},
				[]int{3, 2},
				[]int{4, 5, 6, 7},
				[]int{12, 11, 10, 9, 8},
			},
		},
	}
	ass := require.New(t)
	for _, tc := range testData {
		root := createTree(tc.in, 0)
		ass.Equal(tc.out, zigzagLevelOrder(root))
	}
}

func createTree(arr []int, pos int) *TreeNode {
	length := len(arr) - 1
	if pos > length || arr[pos] == null {
		return nil
	}
	root := new(TreeNode)
	root.Val = arr[pos]
	lch := pos<<1 + 1
	rch := lch + 1
	if lch <= length {
		root.Left = createTree(arr, lch)
	}
	if rch <= length {
		root.Right = createTree(arr, rch)
	}
	return root
}

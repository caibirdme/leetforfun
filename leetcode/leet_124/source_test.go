package leet_124

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const null = -9999999999

func TestMaxPathSum(t *testing.T) {
	var testData = []struct {
		in  []int
		out int
	}{
		{
			in:  []int{-3},
			out: -3,
		},
		{
			in:  []int{1, -2, -3, 1, 3, -2, null, -1},
			out: 3,
		},
		{
			in:  []int{-1, 2, 3, 1, -3, 4, -2, null, null, null, null, null, null, 7},
			out: 12,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		root := &TreeNode{
			Val: tc.in[0],
		}
		buildTree(root, tc.in, len(tc.in), 0)
		ass.Equal(tc.out, maxPathSum(root), "%v", tc.in)
	}
}

func buildTree(root *TreeNode, arr []int, length, idx int) {
	lch := idx<<1 + 1
	if lch >= length {
		return
	}
	if arr[lch] != null {
		tree := new(TreeNode)
		tree.Val = arr[lch]
		buildTree(tree, arr, length, lch)
		root.Left = tree
	}
	rch := lch + 1
	if rch >= length {
		return
	}
	if arr[rch] == null {
		return
	}
	tree := new(TreeNode)
	tree.Val = arr[rch]
	buildTree(tree, arr, length, rch)
	root.Right = tree
}

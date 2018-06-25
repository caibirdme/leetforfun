package leet_572

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

const null = -999999

func TestIsSubtree(t *testing.T) {
	var testCase = []struct {
		s      []int
		t      []int
		expect bool
	}{
		{
			s:      []int{3, 4, 5, 1, 2, null, null},
			t:      []int{4, 1, 2},
			expect: true,
		},
		{
			s:      []int{3, 4, 5, 1, 2, null, null, null, null, 0},
			t:      []int{4, 1, 2},
			expect: false,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		ts := createTree(tc.s, 0)
		tt := createTree(tc.t, 0)
		actual := isSubtree(ts, tt)
		should.Equal(tc.expect, actual, "case#%d failed", idx)
	}
}

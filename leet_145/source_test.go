package leet_145

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const null = 999999

func TestPostorderTraversal(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums: []int{1, null, 2, 3},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, null, 1, 8, 7, 5, 3, null, 10, 2, 7},
		},
		{
			nums: []int{8, 3, 5, 4, null, 2, 9, 10, 11, 7, 6, 12, 20, null, null, 21, 22, null, null, 10, 11},
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		root := createTree(tc.nums)
		actual := postorderTraversal(root)
		var expect []int
		recursiveTraversal(root, &expect)
		ass.Equal(expect, actual, "nums: %v", tc.nums)
	}
}

func recursiveTraversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		(*ans) = append(*ans, root.Val)
		return
	}
	recursiveTraversal(root.Left, ans)
	recursiveTraversal(root.Right, ans)
	(*ans) = append(*ans, root.Val)
}

func createTree(nums []int) *TreeNode {
	length := len(nums)
	if 0 == length {
		return nil
	}
	q := new(list)
	root := &TreeNode{Val: nums[0]}
	q.Push(root)
	for i := 1; !q.Empty() && i < length; i++ {
		node := q.Pop()
		if nums[i] != null {
			node.Left = &TreeNode{
				Val: nums[i],
			}
			q.Push(node.Left)
		}
		i++
		if i < length && nums[i] != null {
			node.Right = &TreeNode{
				Val: nums[i],
			}
			q.Push(node.Right)
		}
	}
	return root
}

type list struct {
	buf []*TreeNode
}

func (l *list) Pop() *TreeNode {
	if len(l.buf) == 0 {
		return nil
	}
	t := l.buf[0]
	l.buf = l.buf[1:]
	return t
}

func (l *list) Push(r *TreeNode) {
	l.buf = append(l.buf, r)
}

func (l *list) Empty() bool {
	return len(l.buf) == 0
}

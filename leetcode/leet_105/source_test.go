package leet_105

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTree(t *testing.T) {
	var testCase = []struct {
		preorder, inorder []int
		expect            []int
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expect:   []int{3, 9, 20, null, null, 15, 7},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := buildTree(tc.preorder, tc.inorder)
		expect := createTree(tc.expect)
		should.True(treeEq(expect, actual), "case #%d fail", idx)
	}
}

func treeEq(a, b *TreeNode) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return treeEq(a.Left, b.Left) && treeEq(a.Right, b.Right)
}

const null = 999999

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

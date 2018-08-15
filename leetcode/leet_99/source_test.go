package leet_99

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecoverTree(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 3, null, null, 2},
			expect: []int{3, 1, null, null, 2},
		},
		{
			nums:   []int{5, 1, 4, null, null, 3, 6},
			expect: []int{3, 1, 4, null, null, 5, 6},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		root := createTree(tc.nums)
		recoverTree(root)
		expect := createTree(tc.expect)
		should.True(treeEq(root, expect), "case #%d fail", idx)
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

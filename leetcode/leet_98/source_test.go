package leet_98

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidBST(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect bool
	}{
		{
			nums:   []int{2, 1, 3},
			expect: true,
		},
		{
			nums:   []int{5, 1, 4, null, null, 3, 6},
			expect: false,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		root := createTree(tc.nums)
		actual := isValidBST(root)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
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

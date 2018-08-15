package leet_114

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlatten(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, null, 3},
			expect: []int{1, 2, 3},
		},
		{
			nums:   []int{2, 1, null, 3, 4},
			expect: []int{2, 1, 3, 4},
		},
		{
			nums:   []int{2, null, 1, 5, 8, null, null, 9, null, null, 6},
			expect: []int{2, 1, 5, 8, 9, 6},
		},
		{
			nums:   []int{1, 2, 5, 3, 4, null, 6},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		root := createTree(tc.nums)
		flatten(root)
		i := 0
		for root != nil {
			should.True(i < len(tc.expect), "case #%d fail", idx)
			should.Nil(root.Left, "case #%d fail", idx)
			should.Equal(tc.expect[i], root.Val, "case #%d fail", idx)
			root = root.Right
			i++
		}
		should.Equal(len(tc.expect), i, "case #%d fail", idx)
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

package leet_450

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteNode(t *testing.T) {
	var testCase = []struct {
		nums1  []int
		key    int
		expect []int
	}{
		{
			nums1:  []int{20, 6, 50, 4, 9, null, null, null, null, 7, 10, null, 8},
			key:    6,
			expect: []int{20, 7, 50, 4, 9, null, null, null, null, 8, 10},
		},
		{
			nums1:  []int{5, 3, 6, 2, 4, null, 7},
			key:    3,
			expect: []int{5, 4, 6, 2, null, null, 7},
		},
		{
			nums1:  []int{5, 3, 6, 2, 4, null, 7},
			key:    7,
			expect: []int{5, 3, 6, 2, 4},
		},
		{
			nums1:  []int{1, 2, 3},
			key:    1,
			expect: []int{3, 2, null},
		},
		{
			nums1:  []int{1},
			key:    1,
			expect: []int{},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		tree1 := createTree(tc.nums1)
		expectTree := createTree(tc.expect)
		should.Equal(expectTree, deleteNode(tree1, tc.key), "case #%d fail", idx)
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

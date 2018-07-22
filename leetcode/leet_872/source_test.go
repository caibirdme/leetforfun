package leet_872

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeafSimilar(t *testing.T) {
	var testCase = []struct {
		nums1, nums2 []int
		expect       bool
	}{
		{
			nums1:  []int{1, 2, 3, null, 4, 5, null, null, null, 6},
			nums2:  []int{8, 9, 7, 4, 5, 6, null},
			expect: false,
		},
		{
			nums1:  []int{1, 2, 3, null, 4, 5, null, null, null, 6},
			nums2:  []int{8, 9, 7, 4, null, 6, null},
			expect: true,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		root1 := createTree(tc.nums1)
		root2 := createTree(tc.nums2)
		actual := leafSimilar(root1, root2)
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

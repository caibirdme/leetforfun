package leet_95

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateTrees(t *testing.T) {
	var testCase = []struct {
		n      int
		expect [][]int
	}{
		{
			n: 1,
			expect: [][]int{
				[]int{1},
			},
		},
		{
			n: 2,
			expect: [][]int{
				[]int{1, null, 2},
				[]int{2, 1},
			},
		},
		{
			n: 3,
			expect: [][]int{
				[]int{1, null, 3, 2},
				[]int{3, 2, null, 1},
				[]int{3, 1, null, null, 2},
				[]int{2, 1, 3},
				[]int{1, null, 2, null, 3},
			},
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := generateTrees(tc.n)
		should.Equal(len(tc.expect), len(actual), "n: %d", tc.n)
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

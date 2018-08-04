package leet_61

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateRight(t *testing.T) {
	var testCase = []struct {
		nums   []int
		k      int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4, 5},
			k:      2,
			expect: []int{4, 5, 1, 2, 3},
		},
		{
			nums:   []int{0, 1, 2},
			k:      4,
			expect: []int{2, 0, 1},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		list := newList(tc.nums)
		actual := rotateRight(list, tc.k)
		should.True(listEqual(actual, newList(tc.expect)), "case #%d fail", idx)
	}
}

func listEqual(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	for l1 != nil {
		if l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		return false
	}
	return true
}

func newList(nums []int) *ListNode {
	head := new(ListNode)
	head.Val = nums[0]
	pre := head
	for i := 1; i < len(nums); i++ {
		node := new(ListNode)
		node.Val = nums[i]
		pre.Next = node
		pre = node
	}
	return head
}

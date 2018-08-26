package leet_143

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReorderList(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			expect: []int{1, 4, 2, 3},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			expect: []int{1, 5, 2, 4, 3},
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		head := createList(tc.nums)
		reorderList(head)
		expect := createList(tc.expect)
		should.True(cmpList(head, expect), "%v", tc.nums)
	}
}

func createList(arr []int) *ListNode {
	head := new(ListNode)
	head.Val = arr[0]
	last := head
	for i := 1; i < len(arr); i++ {
		current := new(ListNode)
		current.Val = arr[i]
		last.Next = current
		last = current
	}
	return head
}

func cmpList(a, b *ListNode) bool {
	currentA, currentB := a, b
	for {
		if currentA == nil && currentB != nil {
			return false
		}
		if currentB == nil && currentA != nil {
			return false
		}
		if currentA == nil && currentB == nil {
			return true
		}
		if currentA.Val != currentB.Val {
			return false
		}
		currentA = currentA.Next
		currentB = currentB.Next
	}
}

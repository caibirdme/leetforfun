package leet_82

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteDuplicates(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 2, 3, 3},
			expect: []int{1},
		},
		{
			nums:   []int{1, 2, 2, 2, 2, 2, 2, 2},
			expect: []int{1},
		},
		{
			nums:   []int{},
			expect: []int{},
		},
		{
			nums:   []int{1},
			expect: []int{1},
		},
		{
			nums:   []int{1, 1},
			expect: []int{},
		},
		{
			nums:   []int{1, 1, 2, 3, 3, 4},
			expect: []int{2, 4},
		},
		{
			nums:   []int{1, 1, 2},
			expect: []int{2},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		input := createList(tc.nums)
		should.Equal(createList(tc.expect), deleteDuplicates(input), "case #%d fail", idx)
	}
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
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

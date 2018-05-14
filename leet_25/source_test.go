package leet_25

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseKGroup(t *testing.T) {
	var testData = []struct {
		in  []int
		k   int
		out []int
	}{
		{
			in:  []int{1, 2, 3, 4, 5},
			k:   2,
			out: []int{2, 1, 4, 3, 5},
		},
		{
			in:  []int{1, 2, 3, 4, 5},
			k:   3,
			out: []int{3, 2, 1, 4, 5},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6},
			k:   3,
			out: []int{3, 2, 1, 6, 5, 4},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6},
			k:   1,
			out: []int{1, 2, 3, 4, 5, 6},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6},
			k:   2,
			out: []int{2, 1, 4, 3, 6, 5},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6},
			k:   4,
			out: []int{4, 3, 2, 1, 5, 6},
		},
		{
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:   3,
			out: []int{3, 2, 1, 6, 5, 4, 7, 8},
		},
	}
	ass := require.New(t)
	for _, tc := range testData {
		list := createList(tc.in)
		actual := reverseKGroup(list, tc.k)
		ass.True(cmpList(actual, createList(tc.out)))
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

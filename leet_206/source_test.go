package leet_206

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseList(t *testing.T) {
	var testData = []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{10, 8, 6, 0},
			out: []int{0, 6, 8, 10},
		},
		{
			in:  []int{0, 1, 0, 1, 0, 0, 1},
			out: []int{1, 0, 0, 1, 0, 1, 0},
		},
		{
			in:  []int{1, 2, 3, 4, 5},
			out: []int{5, 4, 3, 2, 1},
		},
	}
	ass := require.New(t)
	for _, tc := range testData {
		list := createList(tc.in)
		actual := reverseList(list)
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

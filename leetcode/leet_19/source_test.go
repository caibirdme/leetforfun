package leet_19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	inArr []int
	inN   int
	out   []int
}

func TestRemoveNthFromEnd(t *testing.T) {
	var data = []Case{
		Case{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
		Case{
			[]int{1, 2, 3, 4, 5, 6},
			6,
			[]int{2, 3, 4, 5, 6},
		},
		Case{
			[]int{1, 2},
			1,
			[]int{1},
		},
		Case{
			[]int{1},
			1,
			nil,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		head := createNode(tc.inArr)
		now := removeNthFromEnd(head, tc.inN)
		ass.Equal(tc.out, now.output(), "in: %v, n:%d", tc.inArr, tc.inN)
	}
}

func createNode(arr []int) *ListNode {
	if nil == arr || 0 == len(arr) {
		return nil
	}
	head := new(ListNode)
	head.Val = arr[0]
	tail := head
	for i := 1; i < len(arr); i++ {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = arr[i]
	}
	return head
}

package leet_24

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  []int
	out []int
}

func TestSwapPairs(t *testing.T) {
	var data = []Case{
		Case{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 4, 3},
		},
		Case{
			[]int{2},
			[]int{2},
		},
		Case{
			[]int{8, 0},
			[]int{0, 8},
		},
		Case{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
		},
		Case{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 1, 4, 3, 5},
		},
		Case{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{2, 1, 4, 3, 6, 5},
		},
		Case{
			nil,
			[]int{},
		},
		Case{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			[]int{3, 1, 7, 5, 2, 9, 6, 4, 10, 8},
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, output(swapPairs(createNode(tc.in))), "in: %v", tc.in)
	}
}

func createNode(arr []int) *ListNode {
	if nil == arr || len(arr) == 0 {
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

func output(l *ListNode) []int {
	arr := make([]int, 0)
	for l != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return arr
}

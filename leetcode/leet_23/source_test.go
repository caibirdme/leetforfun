package leet_23

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  [][]int
	out []int
}

func TestMergeKLists(t *testing.T) {
	var data = []Case{
		Case{
			[][]int{
				[]int{5, 9},
				[]int{1, 6, 10},
				[]int{3, 4, 5, 7, 8},
				[]int{0, 2, 11},
			},
			[]int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10, 11},
		},
		Case{
			[][]int{
				[]int{0, 0},
				[]int{0, 0},
			},
			[]int{0, 0, 0, 0},
		},
		Case{
			[][]int{
				[]int{1, 4, 7, 10},
				[]int{2, 5, 8, 11},
				[]int{3, 6, 9, 12},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		Case{
			[][]int{
				[]int{0},
			},
			[]int{0},
		},
		Case{
			[][]int{},
			[]int{},
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		var l []*ListNode
		for _, item := range tc.in {
			n := createNode(item)
			ass.Equal(item, output(n), "create in : %v", item)
			l = append(l, n)
		}
		ass.Equal(tc.out, output(mergeKLists(l)), "in: %v", tc.in)
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

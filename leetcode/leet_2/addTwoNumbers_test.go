package leet_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	arr1, arr2, ans []int
}

func TestAddTwoNumbers(t *testing.T) {
	var data = []testCase{
		testCase{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		testCase{
			[]int{1, 2, 3},
			[]int{0},
			[]int{1, 2, 3},
		},
		testCase{
			[]int{7, 8, 9, 9},
			[]int{2, 5, 8, 3, 4, 6},
			[]int{9, 3, 8, 3, 5, 6},
		},
		testCase{
			[]int{1},
			[]int{9, 9},
			[]int{0, 0, 1},
		},
		testCase{
			[]int{2},
			[]int{9, 9, 9, 9, 9, 9},
			[]int{1, 0, 0, 0, 0, 0, 1},
		},
	}
	ass := assert.New(t)
	for _, item := range data {
		l1 := createListNode(item.arr1)
		l2 := createListNode(item.arr2)
		ans := createListNode(item.ans)
		ass.True(eq(ans, addTwoNumbers(l1, l2)))
	}
}

func createListNode(arr []int) *ListNode {
	head := new(ListNode)
	cur := head
	length := len(arr) - 1
	for i := 0; i <= length; i++ {
		cur.Val = arr[i]
		if i == length {
			break
		}
		cur.Next = new(ListNode)
		cur = cur.Next
	}
	return head
}

func eq(l *ListNode, ans *ListNode) bool {
	for ans != nil {
		if l == nil || ans.Val != l.Val {
			return false
		}
		ans = ans.Next
		l = l.Next
	}
	return true
}

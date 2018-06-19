package leet_19

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) output() []int {
	if nil == l {
		return nil
	}
	var arr []int
	t := l
	for nil != t {
		arr = append(arr, t.Val)
		t = t.Next
	}
	return arr
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := head
	for i := 1; i < n; i++ {
		tail = tail.Next
	}
	if nil == tail.Next {
		return head.Next
	}
	current := head
	for tail.Next.Next != nil {
		current = current.Next
		tail = tail.Next
	}
	current.Next = current.Next.Next
	return head
}

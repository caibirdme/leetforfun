package leet_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var pre *ListNode
	current := head
	for current.Next != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	current.Next = pre
	return current
}

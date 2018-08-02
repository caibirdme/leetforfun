package leet_876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	preHead := &ListNode{
		Next: head,
	}
	fast := preHead
	slow := preHead
	over := false
	for !over {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			over = true
		}
		if fast.Next != nil {
			fast = fast.Next
		} else {
			over = true
		}
		slow = slow.Next
	}
	return slow
}

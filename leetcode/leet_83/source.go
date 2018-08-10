package leet_83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := head
	for pre.Next != nil {
		cur := pre.Next
		for cur != nil && cur.Val == pre.Val {
			pre.Next = cur.Next
			cur = cur.Next
		}
		if pre.Next == nil {
			break
		}
		pre = pre.Next
	}
	return head
}

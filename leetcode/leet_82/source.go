package leet_82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fakeHead := &ListNode{Next: head, Val: head.Val - 1}
	pre := fakeHead
	for pre != nil {
		cur := pre.Next
		if cur == nil {
			break
		}
		dump := false
		if cur.Next != nil && cur.Next.Val == cur.Val {
			dump = true
		}
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		if dump {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}
	return fakeHead.Next
}

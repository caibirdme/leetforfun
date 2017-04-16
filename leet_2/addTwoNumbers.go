package leet_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cpl1 := l1
	cpl2 := l2
	for cpl1 != nil && cpl2 != nil {
		cpl1 = cpl1.Next
		cpl2 = cpl2.Next
	}
	if cpl1 == nil {
		l1, l2 = l2, l1
	}
	cpl1 = l1
	cpl2 = l2
	for cpl2 != nil {
		cpl1.Val += cpl2.Val
		if cpl1.Val >= 10 {
			if cpl1.Next == nil {
				cpl1.Next = new(ListNode)
			}
			cpl1.Next.Val++
			cpl1.Val -= 10
		}
		cpl1 = cpl1.Next
		cpl2 = cpl2.Next
	}
	for cpl1 != nil && cpl1.Val >= 10 {
		if cpl1.Next == nil {
			cpl1.Next = new(ListNode)
		}
		cpl1.Next.Val++
		cpl1.Val -= 10
		cpl1 = cpl1.Next
	}
	return l1
}

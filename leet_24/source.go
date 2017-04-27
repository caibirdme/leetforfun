package leet_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	one, two, three := head, head.Next, head.Next.Next
	two.Next, one.Next = one, swapPairs(three)
	return two
}

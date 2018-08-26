package leet_143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var reverseHead *reverseList
	var n int
	reverseListNode(head, &reverseHead, &n)
	n >>= 1
	cur := head
	tail := reverseHead
	for i := 0; i < n; i++ {
		if cur.Next == tail.Val {
			break
		}
		t := cur.Next
		tail.Pre.Val.Next = nil
		cur.Next = tail.Val
		cur.Next.Next = t
		tail = tail.Pre
		cur = t
	}
}

func reverseListNode(cur *ListNode, reverseHead **reverseList, n *int) *reverseList {
	*n++
	if cur.Next == nil {
		*reverseHead = &reverseList{Val: cur}
		return *reverseHead
	}
	pre := reverseListNode(cur.Next, reverseHead, n)
	this := &reverseList{Val: cur}
	pre.Pre = this
	return this
}

type reverseList struct {
	Pre *reverseList
	Val *ListNode
}

package leet_61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	cur := head
	count := 1
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	cur.Next = head
	k %= count
	cur = head
	length := count - k
	for i := 1; i < length; i++ {
		cur = cur.Next
	}
	ans := cur.Next
	cur.Next = nil
	return ans
}

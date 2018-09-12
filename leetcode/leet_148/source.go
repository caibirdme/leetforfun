package leet_148

import (
	"github.com/caibirdme/leetforfun/testsuit"
)

type ListNode = testsuit.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	nH, _ := mergeSort(head, cur, count)
	return nH
}

func mergeSort(head, tail *ListNode, n int) (*ListNode, *ListNode) {
	if n == 1 {
		return head, head
	}
	mid := n / 2
	count := 1
	preTail := head
	for count < mid {
		preTail = preTail.Next
		count++
	}
	postHead := preTail.Next
	preTail.Next = nil
	var preHead, postTail *ListNode
	preHead, preTail = mergeSort(head, preTail, mid)
	postHead, postTail = mergeSort(postHead, tail, n-mid)
	postTail.Next = nil
	if postHead.Val < preHead.Val {
		postHead, preHead = preHead, postHead
		postTail, preTail = preTail, postTail
	}
	head = preHead
	pre := head
	preHead = preHead.Next
	for preHead != nil && postHead != nil {
		if preHead.Val <= postHead.Val {
			pre = preHead
			preHead = preHead.Next
			continue
		}
		t := postHead
		postHead = postHead.Next
		pre.Next = t
		t.Next = preHead
		pre = t
	}
	if postHead != nil {
		pre.Next = postHead
		tail = postTail
	} else {
		tail = preTail
	}
	return head, tail
}

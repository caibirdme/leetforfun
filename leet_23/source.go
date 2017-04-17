package leet_23

import (
	"container/heap"
)

type H []*ListNode

func (h H) Len() int {
	return len(h)
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h H) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *H) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *H) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	q := make(H, 0)
	for i := 0; i < len(lists); i++ {
		if nil != lists[i] {
			q = append(q, lists[i])
		}
	}
	if q.Len() == 0 {
		return nil
	}
	heap.Init(&q)
	head := new(ListNode)
	tail := head
	for q.Len() > 0 {
		e := heap.Pop(&q).(*ListNode)
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = e.Val
		if nil != e.Next {
			e = e.Next
			heap.Push(&q, e)
		}
	}
	return head.Next
}

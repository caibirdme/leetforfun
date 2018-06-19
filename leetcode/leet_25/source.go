package leet_25

type ListNode struct {
	Val  int
	Next *ListNode
}

type dequeue struct {
	buff []*ListNode
}

func (q *dequeue) lpop() *ListNode {
	if 0 == len(q.buff) {
		return nil
	}
	p := q.buff[0]
	q.buff = q.buff[1:]
	return p
}

func (q *dequeue) rpop() *ListNode {
	if 0 == len(q.buff) {
		return nil
	}
	last := len(q.buff) - 1
	p := q.buff[last]
	q.buff = q.buff[:last]
	return p
}

func (q *dequeue) rpush(p *ListNode) {
	q.buff = append(q.buff, p)
}

func (q *dequeue) clear() {
	q.buff = buff[:0]
}

var buff []*ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k <= 1 {
		return head
	}
	buff = make([]*ListNode, 0, k)
	queue := &dequeue{buff: buff}
	count := 0
	cur := head
	for cur != nil {
		count++
		queue.rpush(cur)
		if count == k {
			count = 0
			for {
				l := queue.lpop()
				r := queue.rpop()
				if l == nil || r == nil {
					break
				}
				l.Val, r.Val = r.Val, l.Val
			}
			queue.clear()
		}
		cur = cur.Next
	}
	return head
}

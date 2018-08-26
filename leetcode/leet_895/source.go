package leet_895

import (
	"container/heap"
)

type FreqStack struct {
	Top       int
	PriorityQ *MyHeap
	Index     map[int]*MapElem
}

func Constructor() FreqStack {
	var pq MyHeap
	stack := FreqStack{
		Top:       -1,
		PriorityQ: &pq,
		Index:     make(map[int]*MapElem),
	}
	heap.Init(stack.PriorityQ)
	return stack
}

func (this *FreqStack) Push(x int) {
	this.Top++
	if _, ok := this.Index[x]; !ok {
		this.Index[x] = &MapElem{
			Pos: []int{this.Top},
		}
	} else {
		this.Index[x].Pos = append(this.Index[x].Pos, this.Top)
	}
	info := this.Index[x]
	if info.Elem != nil {
		info.Elem.Freq++
		info.Elem.Pos = info.GetTop()
		heap.Fix(this.PriorityQ, info.Elem.Idx)
	} else {
		hElem := &HeapElem{
			Val:  x,
			Freq: len(info.Pos),
			Pos:  info.GetTop(),
		}
		this.Index[x].Elem = hElem
		heap.Push(this.PriorityQ, hElem)
	}

}

func (this *FreqStack) Pop() int {
	v := (*this.PriorityQ)[0].Val
	(*this.PriorityQ)[0].Freq--
	if (*this.PriorityQ)[0].Freq == 0 {
		heap.Remove(this.PriorityQ, 0)
		delete(this.Index, v)
	} else {
		this.Index[v].Pos = this.Index[v].Pos[:len(this.Index[v].Pos)-1]
		this.Index[v].Elem.Pos = this.Index[v].GetTop()
		heap.Fix(this.PriorityQ, 0)
	}
	return v
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */

type MapElem struct {
	Pos  []int
	Elem *HeapElem
}

func (m *MapElem) GetTop() int {
	if len(m.Pos) == 0 {
		return 0
	}
	return m.Pos[len(m.Pos)-1]
}

type HeapElem struct {
	Val  int
	Freq int
	Pos  int
	Idx  int
}

type MyHeap []*HeapElem

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	if h[i].Freq > h[j].Freq {
		return true
	}
	if h[j].Freq == h[i].Freq {
		return h[i].Pos > h[j].Pos
	}
	return false
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Idx = i
	h[j].Idx = j
}

func (h *MyHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*HeapElem)
	item.Idx = n
	*h = append(*h, item)
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.Idx = -1 // for safety
	*h = old[0 : n-1]
	return item
}

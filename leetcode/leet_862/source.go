package leet_862

import (
	"container/list"
)

func shortestSubarray(A []int, K int) int {
	length := len(A)
	if length == 1 {
		if A[0] >= K {
			return 1
		}
		return -1
	}
	pre := 0
	sum := make([]int, length+1)
	for i := 0; i < length; i++ {
		pre += A[i]
		sum[i+1] = pre
	}
	l := list.New()
	l.PushBack(0)
	ans := length + 1
	for i := 1; i <= length; i++ {
		for l.Len() > 0 && sum[i] <= sum[l.Back().Value.(int)] {
			l.Remove(l.Back())
		}
		for l.Len() > 0 && sum[i]-K >= sum[l.Front().Value.(int)] {
			front := l.Front()
			ans = min(ans, i-front.Value.(int))
			l.Remove(front)
		}
		l.PushBack(i)
	}
	if ans == length+1 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

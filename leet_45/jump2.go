package leet_45

import (
	"container/list"
)

// build graph and use SPFA 2 find the shortest path
func jump(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	dist := make([]int, length)
	last := length - 1
	for i := 1; i < length; i++ {
		dist[i] = length
	}
	l := list.New()
	l.PushBack(0)
	for l.Len() > 0 {
		idx := l.Remove(l.Front()).(int)
		for i := 1; i <= nums[idx]; i++ {
			t := idx + i
			if t == last {
				return dist[idx] + 1
			}
			n := dist[idx] + 1
			if dist[t] <= n {
				continue
			}
			dist[t] = n
			if nums[t] > 0 {
				l.PushBack(t)
			}
		}
	}
	return dist[last]
}

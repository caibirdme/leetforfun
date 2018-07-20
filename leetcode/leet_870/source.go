package leet_870

import "sort"

type desc []int

func (d desc) Len() int {
	return len(d)
}

func (d desc) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d desc) Less(i, j int) bool {
	return d[i] > d[j]
}

func advantageCount(A []int, B []int) []int {
	length := len(B)
	newB := make([]int, length)
	copy(newB, B)
	sort.Sort(desc(A))
	sort.Sort(desc(newB))
	left := 0
	n := length - 1
	right := n
	ans := make([]int, length)
	for left <= right {
		bIdx := left + n - right
		if A[left] > newB[bIdx] {
			ans[bIdx] = A[left]
			left++
		} else {
			ans[bIdx] = A[right]
			right--
		}
	}
	mp := make(map[int][]int)
	for idx, v := range newB {
		if _, ok := mp[v]; !ok {
			mp[v] = []int{ans[idx]}
		} else {
			mp[v] = append(mp[v], ans[idx])
		}
	}
	for idx, v := range B {
		ans[idx] = mp[v][0]
		if len(mp[v]) > 1 {
			mp[v] = mp[v][1:]
		}
	}
	return ans
}

func bigger(a, b int) bool {
	return a > b
}

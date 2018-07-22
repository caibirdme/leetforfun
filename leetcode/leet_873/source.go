package leet_873

import (
	"sort"
)

func lenLongestFibSubseq(A []int) int {
	length := len(A)
	ans := 0
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			pp, p := A[i], A[j]
			count := 2
			t := pp + p
			for find(A, t) {
				pp = p
				p = t
				t = pp + p
				count++
				ans = max(ans, count)
			}
		}
	}
	return ans
}

func find(A []int, n int) bool {
	idx := sort.SearchInts(A, n)
	if len(A) > idx && A[idx] == n {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

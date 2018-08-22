package leet_891

import (
	"sort"
)

func sumSubseqWidths(A []int) int {
	const mod = 1000000007
	sort.Ints(A)
	length := len(A)
	f := make([]int, length)
	f[0] = 1
	for i := 1; i < length; i++ {
		f[i] = f[i-1] * 2
		if f[i] >= mod {
			f[i] %= mod
		}
	}
	last := length - 1
	ans := 0
	for i := 0; i < length; i++ {
		ans += ((f[i] - f[last-i]) * A[i]) % mod
		ans %= mod
	}
	return ans
}

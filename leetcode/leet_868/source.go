package leet_868

func binaryGap(N int) int {
	pre := -1
	ans := 0
	idx := 0
	for N > 0 {
		b := N & 1
		if b == 1 {
			if pre != -1 {
				ans = max(ans, idx-pre)
			}
			pre = idx
		}
		N >>= 1
		idx++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

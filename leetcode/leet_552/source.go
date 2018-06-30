package leet_552

const delta = 1000000007

func checkRecord(n int) int {
	if n == 1 {
		return 3
	}
	f := make([][3]int, n)
	f[0][0] = 1
	f[0][1] = 1
	f[1][0] = 2
	f[1][1] = 1
	f[1][2] = 1
	// 把A放到第一个位置或者最后一个位置 + 完全不放A
	ans := ((getN(n-2, f)<<1)%delta + getN(n-1, f)) % delta
	// 枚举每个位置，放A
	for i := 1; i < n-1; i++ {
		// 左边还有left个格子
		left := i - 1
		// 右边还有right个格子
		right := n - 2 - i
		ans += (getN(left, f) * getN(right, f)) % delta
		ans %= delta
	}
	return ans
}

func getN(n int, f [][3]int) int {
	return (search(n, 0, f) + search(n, 1, f) + search(n, 2, f)) % delta
}

func search(n int, k int, f [][3]int) int {
	if n == 0 {
		return f[0][k]
	}
	if f[n][k] > 0 {
		return f[n][k]
	}
	var ans int
	switch k {
	case 0:
		ans = search(n-1, 0, f) + search(n-1, 1, f) + search(n-1, 2, f)
	case 1:
		ans = search(n-1, 0, f)
	case 2:
		ans = search(n-1, 1, f)
	}
	ans %= delta
	f[n][k] = ans
	return ans
}

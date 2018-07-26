package leet_115

func numDistinct(s string, t string) int {
	lens := len(s)
	lent := len(t)
	if lent == 0 {
		return 0
	}
	if lens <= lent {
		if s == t {
			return 1
		}
		return 0
	}
	var f = [2][]int{
		make([]int, lens),
		make([]int, lens),
	}
	for i := 0; i < lens; i++ {
		if s[i] == t[0] {
			f[0][i] = 1
		}
	}
	idx := 1
	sum := make([]int, lens)
	for pre, i := 0, 0; i < lens; i++ {
		sum[i] = pre + f[0][i]
		pre = sum[i]
	}
	for i := 1; i < lent; i++ {
		for j := 0; j < i; j++ {
			f[idx][j] = 0
		}
		for j := i; j < lens; j++ {
			if t[i] == s[j] {
				f[idx][j] = sum[j-1]
			} else {
				f[idx][j] = 0
			}
		}
		for pre, j := 0, 0; j < lens; j++ {
			sum[j] = pre + f[idx][j]
			pre = sum[j]
		}
		idx = 1 - idx
	}
	return sum[lens-1]
}

// 只用一维数组
func numDistinct2(s string, t string) int {
	lens := len(s)
	lent := len(t)
	if lent == 0 {
		return 0
	}
	if lens <= lent {
		if s == t {
			return 1
		}
		return 0
	}
	var f = make([]int, lens)
	for i := 0; i < lens; i++ {
		if s[i] == t[0] {
			f[i] = 1
		}
	}
	sum := make([]int, lens)
	for pre, i := 0, 0; i < lens; i++ {
		sum[i] = pre + f[i]
		pre = sum[i]
	}
	for i := 1; i < lent; i++ {
		for j := 0; j < i; j++ {
			f[j] = 0
		}
		for j := lens - 1; j >= i; j-- {
			if t[i] == s[j] {
				f[j] = sum[j-1]
			} else {
				f[j] = 0
			}
		}
		for pre, j := 0, 0; j < lens; j++ {
			sum[j] = pre + f[j]
			pre = sum[j]
		}
	}
	return sum[lens-1]
}

package leet_96

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			left := j - 1
			right := i - j
			f[i] += f[left] * f[right]
		}
	}
	return f[n]
}

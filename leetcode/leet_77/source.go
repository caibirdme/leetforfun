package leet_77

func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return nil
	}
	var ans [][]int
	num := make([]int, k)
	dfs(0, 0, k, n, num, &ans)
	return ans
}

func dfs(i, j, k, n int, num []int, ans *[][]int) {
	if j == k {
		*ans = append(*ans, append([]int(nil), num[:k]...))
		return
	}
	rest := k - j - 1
	maxN := n - rest
	for t := i + 1; t <= maxN; t++ {
		num[j] = t
		dfs(t, j+1, k, n, num, ans)
	}
}

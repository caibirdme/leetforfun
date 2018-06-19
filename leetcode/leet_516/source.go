package leet_516

func longestPalindromeSubseq(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	var mem [1001][1001]int
	return dfs([]byte(s), 0, length-1, &mem)
}

func dfs(arr []byte, i, j int, mem *[1001][1001]int) int {
	if i == j {
		return 1
	}
	if i+1 == j && arr[i] == arr[j] {
		return 2
	}
	if i > j {
		return 0
	}
	if t := (*mem)[i][j]; t != 0 {
		return t
	}
	ans := 0
	if arr[i] == arr[j] {
		ans = dfs(arr, i+1, j-1, mem) + 2
	}
	ans = max(ans, dfs(arr, i+1, j, mem), dfs(arr, i, j-1, mem))
	(*mem)[i][j] = ans
	return ans
}

func max(a, b, c int) int {
	ans := a
	if ans < b {
		ans = b
	}
	if ans < c {
		ans = c
	}
	return ans
}

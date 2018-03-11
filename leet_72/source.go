package leet_72

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}
	n := len(word1)
	m := len(word2)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			f[i][j] = -1
		}
	}
	has := false
	if word1[0] == word2[0] {
		f[0][0] = 0
		has = true
	} else {
		f[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if word1[0] == word2[i] {
			f[0][i] = i
			has = true
		} else if has {
			f[0][i] = i
		} else {
			f[0][i] = i + 1
		}
	}
	if f[0][0] == 0 {
		has = true
	} else {
		has = false
	}
	for i := 1; i < n; i++ {
		if word2[0] == word1[i] {
			f[i][0] = i
			has = true
		} else if has {
			f[i][0] = i
		} else {
			f[i][0] = i + 1
		}
	}
	return solve([]byte(word1), []byte(word2), n-1, m-1, &f)
}

func solve(word1, word2 []byte, n, m int, mem *[][]int) int {
	if t := (*mem)[n][m]; t != -1 {
		return t
	}
	// answer can't bigger than m+n
	ans := m + n
	if word1[n] == word2[m] {
		ans = solve(word1, word2, n-1, m-1, mem)
	}

	ans = min(
		ans,
		solve(word1, word2, n-1, m-1, mem)+1,
		solve(word1, word2, n, m-1, mem)+1,
		solve(word1, word2, n-1, m, mem)+1,
	)
	(*mem)[n][m] = ans
	return ans
}

func min(nums ...int) int {
	t := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if t > nums[i] {
			t = nums[i]
		}
	}
	return t
}

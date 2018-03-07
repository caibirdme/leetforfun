package leet_60

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	ans := make([]byte, n)
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] * i
	}
	var bitmap uint = (1 << uint(n)) - 1
	dfs(1, k, n, &bitmap, f, &ans)
	return string(ans)
}

func dfs(idx, k, n int, bitmap *uint, cache []int, ans *[]byte) {
	if idx == n {
		(*ans)[idx-1] = getKthElement(bitmap, 1, n)
		return
	}
	t := cache[n-idx]
	sum := t
	var i int
	for i = 1; i <= n && sum < k; i++ {
		sum += t
	}
	(*ans)[idx-1] = getKthElement(bitmap, i, n)
	dfs(idx+1, k-sum+t, n, bitmap, cache, ans)
}

func getKthElement(bitmap *uint, k, n int) byte {
	count := 0
	for i := 0; i <= n; i++ {
		var t uint = 1 << uint(i)
		if *bitmap&t == t {
			count++
			if count == k {
				*bitmap ^= t
				return '0' + byte(i) + 1
			}
		}
	}
	return '0'
}

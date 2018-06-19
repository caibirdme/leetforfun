package leet_5

func longestPalindrome(s string) string {
	var f [1001][1001]bool
	var ans, idx int
	ans = 1
	length := len(s)
	for i := 0; i < length; i++ {
		f[i][1] = true
	}
	for i := 0; i < length-1; i++ {
		if s[i] == s[i+1] {
			f[i][2] = true
			if ans < 2 {
				ans = 2
				idx = i
			}
		}
	}
	for i := 3; i <= length; i++ {
		for j := 0; j <= length-i; j++ {
			if s[j] == s[j+i-1] && f[j+1][i-2] {
				f[j][i] = true
				if i > ans {
					ans = i
					idx = j
				}
			}
		}
	}
	return s[idx : idx+ans]
}

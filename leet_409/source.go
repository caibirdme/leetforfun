package leet_409

func longestPalindrome(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	count := make([]int, 256)
	for _, c := range s {
		count[c]++
	}
	single := false
	ans := 0
	for i := 0; i < 256; i++ {
		c := count[i]
		if c >= 2 {
			t := c / 2 * 2
			ans += t
			c -= t
		}
		if c == 1 {
			single = true
		}
	}
	if single {
		ans++
	}
	return ans
}

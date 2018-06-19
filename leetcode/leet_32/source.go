package leet_32

func longestValidParentheses(s string) int {
	if "" == s {
		return 0
	}
	length := len(s)
	var ans int
	arr := []byte(s)
	f := make([]int, length)
	for i := 1; i < length; i++ {
		if arr[i] == '(' {
			continue
		}
		if arr[i-1] == '(' {
			f[i] = 2
			if i-2 >= 0 {
				f[i] += f[i-2]
			}
		} else if p := i - f[i-1] - 1; f[i-1] != 0 && p >= 0 && arr[p] == '(' {
			f[i] = f[i-1] + 2
			if t := p - 1; t >= 0 {
				f[i] += f[t]
			}
		}
		if f[i] > ans {
			ans = f[i]
		}
	}
	return ans
}

package kmp

// FindSubStr ...
func FindSubStr(str, sub string) int {
	next := getNext([]byte(sub))
	lensub := len(sub)
	s := []byte(str)
	lens := len(s)
	i := 0
	j := 0
	count := 0
	for i < lens {
		if s[i] == sub[j] {
			count++
			if count == lensub {
				return i - count + 1
			}
			i++
			j++
		} else if j != 0 {
			j = next[j] + 1
			count = max(next[j], 0)
		} else {
			count = 0
			j = 0
			i++
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getNext(sub []byte) []int {
	lensub := len(sub)
	next := make([]int, lensub)
	for i := 0; i < lensub; i++ {
		next[i] = -1
	}
	for i := 1; i < lensub; i++ {
		if sub[i] == sub[next[i-1]+1] {
			next[i] = next[i-1] + 1
		} else {
			t := i - 1
			for next[t] != -1 {
				if sub[i] == sub[next[t]+1] {
					next[i] = next[t] + 1
				} else {
					t = next[t]
				}
			}
		}
	}
	return next
}

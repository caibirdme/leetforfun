package leet_76

func minWindow(s string, t string) string {
	target := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if _, ok := target[t[i]]; !ok {
			target[t[i]] = 1
		} else {
			target[t[i]]++
		}
	}
	count := make(map[byte]int)
	ans := len(s) + 1
	pos := make([]int, 0, len(target))
	left := 0
	res := ""
	for i := 0; i < len(s); i++ {
		if _, ok := target[s[i]]; !ok {
			continue
		}
		addCount(count, s[i])
		pos = append(pos, i)
		if mapEq(target, count) {
			l := i - pos[left] + 1
			if ans > l {
				ans = l
				res = s[pos[left] : i+1]
			}
			for {
				ch := s[pos[left]]
				left++
				if left >= len(pos) {
					break
				}
				reduceCount(count, ch)
				if mapEq(target, count) {
					l := i - pos[left] + 1
					if ans > l {
						ans = l
						res = s[pos[left] : i+1]
					}
				} else {
					left--
					addCount(count, ch)
					break
				}
			}
		}
	}
	return res
}

func reduceCount(count map[byte]int, ch byte) {
	v := count[ch]
	if v > 1 {
		count[ch]--
	} else {
		delete(count, ch)
	}
}

func addCount(count map[byte]int, ch byte) {
	if _, ok := count[ch]; !ok {
		count[ch] = 1
	} else {
		count[ch]++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mapEq(mp1, mp2 map[byte]int) bool {
	if len(mp1) != len(mp2) {
		return false
	}
	for k, v1 := range mp1 {
		v2, ok := mp2[k]
		if !ok {
			return false
		}
		if v1 > v2 {
			return false
		}
	}
	return true
}

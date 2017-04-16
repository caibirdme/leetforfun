package leet_3

func lengthOfLongestSubstring(s string) int {
	var hash [256]int
	for i := 0; i < 256; i++ {
		hash[i] = -1
	}
	max := 0
	cur := 0
	for idx := range s {
		c := s[idx]
		pos := hash[c]
		if pos == -1 || idx-pos >= cur+1 {
			cur++
			if max < cur {
				max = cur
			}
		} else {
			cur = idx - pos
		}
		hash[c] = idx
	}
	return max
}

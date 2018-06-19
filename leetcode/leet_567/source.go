package leet_567

func checkInclusion(s1 string, s2 string) bool {
	lens1 := len(s1)
	lens2 := len(s2)
	if lens1 > lens2 {
		return false
	}
	record := make([]int, 26)
	for i := 0; i < lens1; i++ {
		record[s1[i]-'a']++
	}
	accumulate := make([]int, 26)
	for i := 0; i < lens1; i++ {
		accumulate[s2[i]-'a']++
	}
	left := 0
	current := lens1
	if check(record, accumulate) {
		return true
	}
	for ; current < lens2; current++ {
		l := s2[left] - 'a'
		c := s2[current] - 'a'
		accumulate[c]++
		accumulate[l]--
		left++
		if accumulate[c] == record[c] && accumulate[l] == record[l] {
			if check(record, accumulate) {
				return true
			}
		}
	}
	return false
}

func check(record, accumulate []int) bool {
	for i := 0; i < 26; i++ {
		if record[i] != accumulate[i] {
			return false
		}
	}
	return true
}

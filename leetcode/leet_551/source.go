package leet_551

func checkRecord(s string) bool {
	length := len(s)
	countA := 0
	countL := 0
	for i := 0; i < length; i++ {
		if s[i] == 'A' {
			if countA == 1 {
				return false
			}
			countA = 1
			countL = 0
		} else if s[i] == 'L' {
			if countL == 2 {
				return false
			}
			countL++
		} else {
			countL = 0
		}
	}
	return true
}

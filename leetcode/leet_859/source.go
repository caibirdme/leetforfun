package leet_859

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	count := make([]int, 26)
	misMatch := 0
	var x, y int
	multi := false
	for i := 0; i < len(A); i++ {
		t := A[i] - 'a'
		count[t]++
		if count[t] > 1 {
			multi = true
		}
		if A[i] != B[i] {
			misMatch++
			if misMatch > 2 {
				return false
			}
			if misMatch == 1 {
				x = i
			} else {
				y = i
			}
		}
	}
	if misMatch == 0 && !multi {
		return false
	}

	return A[x] == B[y] && B[x] == A[y]
}

package leet_97

func isInterleave(s1 string, s2 string, s3 string) bool {
	lens1 := len(s1)
	lens2 := len(s2)
	lens3 := len(s3)
	if lens3 != lens1+lens2 {
		return false
	}
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	f := make([][]bool, lens1+1)
	for i := 0; i <= lens1; i++ {
		f[i] = make([]bool, lens2+1)
	}
	for i := 1; i <= lens1; i++ {
		t := i - 1
		if s1[t] == s3[t] {
			f[i][0] = true
		} else {
			break
		}
	}
	for i := 1; i <= lens2; i++ {
		t := i - 1
		if s2[t] == s3[t] {
			f[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= lens1; i++ {
		for j := 1; j <= lens2; j++ {
			t := i + j - 1
			f[i][j] = (s1[i-1] == s3[t] && f[i-1][j]) || (s2[j-1] == s3[t] && f[i][j-1])
		}
	}
	return f[lens1][lens2]
}

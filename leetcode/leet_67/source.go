package leet_67

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	lena := len(a)
	lenb := len(b)
	res := make([]byte, lenb+1)
	ba := make([]byte, lena)
	bb := make([]byte, lenb)
	for i := 0; i < lena; i++ {
		ba[i] = a[lena-i-1] - '0'
	}
	for i := 0; i < lenb; i++ {
		bb[i] = b[lenb-i-1] - '0'
	}
	for i := 0; i < lena; i++ {
		res[i] += ba[i] + bb[i]
		if res[i] >= 2 {
			res[i+1]++
			res[i] -= 2
		}
	}
	for i := lena; i < lenb; i++ {
		res[i] += bb[i]
		if res[i] >= 2 {
			res[i] -= 2
			res[i+1]++
		}
	}
	if res[lenb] == 0 {
		lenb--
	}
	i, j := 0, lenb
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	for i := 0; i <= lenb; i++ {
		res[i] += '0'
	}
	return string(res[:lenb+1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

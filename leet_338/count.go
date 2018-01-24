package leet_338

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	f := make([]int, num+1)
	f[0] = 0
	f[1] = 1
	if num == 1 {
		return f
	}
	baseNumber := 2
	exp := 0
	for i := 2; i <= num; i++ {
		if i == baseNumber {
			baseNumber = baseNumber << 1
			exp = (exp << 1) | 1
		}
		f[i] = f[i&exp] + 1
	}
	return f
}

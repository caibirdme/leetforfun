package leet_69

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 1, x>>1
	var ans int
	for l <= r {
		mid := (l + r) >> 1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

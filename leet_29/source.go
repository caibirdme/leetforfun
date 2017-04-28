package leet_29

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	flag := 1
	if dividend < 0 {
		flag = -flag
	}
	if divisor < 0 {
		flag = -flag
	}
	dividend = abs(dividend)
	divisor = abs(divisor)
	if dividend < divisor {
		return 0
	}
	dd := divisor
	step := 1
	var ans int
	for dividend >= dd {
		dividend -= dd
		ans += step
		dd += dd
		step += step
	}
	if dividend >= divisor {
		ans += divide(dividend, divisor)
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if flag < 0 {
		return -ans
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

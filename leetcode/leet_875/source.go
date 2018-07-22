package leet_875

func minEatingSpeed(piles []int, H int) int {
	r := piles[0]
	length := len(piles)
	for i := 1; i < length; i++ {
		r = max(r, piles[i])
	}
	if length == H {
		return r
	}
	ans := r
	l := 1
	for l <= r {
		mid := (l + r) >> 1
		count := H
		ok := true
		for _, num := range piles {
			count -= num / mid
			num %= mid
			if num != 0 {
				count--
			}
			if count < 0 {
				ok = false
				break
			}
		}
		if ok {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

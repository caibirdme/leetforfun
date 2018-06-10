package leet_849

func maxDistToClosest(seats []int) int {
	length := len(seats)
	if length == 2 {
		return 1
	}
	maxCount := 1
	consecutiveZero := 0
	pre := 0
	for i := 0; i < length; i++ {
		if seats[i] == 0 {
			if pre == 0 {
				consecutiveZero++
				maxCount = max(maxCount, consecutiveZero)
			} else {
				consecutiveZero = 1
			}
		} else {
			consecutiveZero = 0
		}
		pre = seats[i]
	}
	ans := max(preZero(seats), tailZero(seats))
	if maxCount&1 == 1 {
		maxCount++
	}
	return max(ans, maxCount>>1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func preZero(seats []int) int {
	count := 0
	length := len(seats)
	for i := 0; i < length; i++ {
		if seats[i] == 0 {
			count++
		} else {
			break
		}
	}
	return count
}

func tailZero(seats []int) int {
	count := 0
	length := len(seats)
	for i := length - 1; i >= 0; i-- {
		if seats[i] == 0 {
			count++
		} else {
			break
		}
	}
	return count
}

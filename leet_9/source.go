package leet_9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	max := 1
	for max <= x {
		max *= 10
	}
	max /= 10
	for x >= 10 {
		if x/max != x%10 {
			return false
		}
		x = ((x - x%10) % max) / 10
		max /= 100
	}
	return true
}

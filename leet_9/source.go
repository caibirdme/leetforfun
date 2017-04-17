package leet_9

func isPalindrome(x int) bool {
	num := x
	var y int
	for num > 0 {
		y = y*10+num%10
		num /= 10
	}
	return x == y
}

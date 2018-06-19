package leet_7

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	num := x
	if x < 0 {
		num = -num
	}
	str := strconv.Itoa(num)
	length := len(str)
	var arr []byte
	for i := length - 1; i >= 0; i-- {
		arr = append(arr, str[i])
	}
	res := string(arr)
	if isOverflow(res) {
		return 0
	}
	if x < 0 {
		res = "-" + res
	}
	ans, _ := strconv.Atoi(res)
	return ans
}

func isOverflow(str string) bool {
	ms := strconv.Itoa(math.MaxInt32)
	if len(str) > len(ms) {
		return true
	} else if len(str) < len(ms) {
		return false
	}
	return str > ms
}

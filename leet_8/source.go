package leet_8

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	ans, _ := convert(str)
	return ans
}

func convert(str string) (int, error) {
	str = strings.Trim(str, " ")
	str = strings.TrimLeft(str, "0")
	length := len(str)
	for i := 0; i < length; i++ {
		if !isDigit(str[i]) && !isOperator(str[i]) {
			str = str[:i]
			break
		}
	}
	if "" == str {
		return 0, nil
	}
	ret, err := strconv.Atoi(str)
	if ret > math.MaxInt32 {
		ret = math.MaxInt32
	} else if ret < -math.MaxInt32 {
		ret = -math.MaxInt32 - 1
	}
	return ret, err
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isOperator(c byte) bool {
	return c == '-' || c == '+'
}

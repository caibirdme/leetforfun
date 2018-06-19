package leet_43

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}
	zeros := 0
	for i := len(num1) - 1; i > 0; i-- {
		if num1[i] == '0' {
			zeros++
		} else {
			break
		}
	}
	num1 = num1[:len(num1)-zeros]
	t := 0
	for i := len(num2) - 1; i > 0; i-- {
		if num2[i] == '0' {
			t++
		} else {
			break
		}
	}
	zeros += t
	num2 = num2[:len(num2)-t]

	mod := 10000
	modL := 4
	a := concat(num1, mod)
	b := concat(num2, mod)
	result := make([]int, len(a)+len(b)+1)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			result[i+j] += a[i] * b[j]
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= mod {
			t := result[i] % mod
			result[i+1] += result[i] / mod
			result[i] = t
		}
	}
	empty := 0
	for i := len(result) - 1; i > 0; i-- {
		if result[i] == 0 {
			empty++
		} else {
			break
		}
	}
	result = result[:len(result)-empty]
	ans := ""
	for i := len(result) - 1; i >= 0; i-- {
		v := strconv.Itoa(result[i])
		for len(v) < modL {
			v = "0" + v
		}
		ans += v
	}
	for zeros > 0 {
		ans += "0"
		zeros--
	}
	preZeros := 0
	for i := 0; i < len(ans); i++ {
		if ans[i] == '0' {
			preZeros++
		} else {
			break
		}
	}
	return ans[preZeros:]
}

func concat(num string, mod int) []int {
	length := len(num)
	step := -1
	for mod > 0 {
		step++
		mod /= 10
	}
	result := make([]int, 0, len(num)/step+1)
	temp := make([]byte, 0, step)
	for i := length - 1; i >= 0; i-- {
		temp = append(temp, num[i])
		if len(temp) == step {
			result = append(result, byte2Int(temp))
			temp = temp[:0]
		}
	}
	if len(temp) > 0 {
		result = append(result, byte2Int(temp))
	}
	return result
}

func byte2Int(arr []byte) int {
	var result int
	base := 1
	for _, b := range arr {
		result += base * int(b-'0')
		base *= 10
	}
	return result
}

func reverse(str string) []byte {
	length := len(str)
	result := make([]byte, length)
	pos := length - 1
	for i := pos; i >= 0; i-- {
		result[pos-i] = str[i]
	}
	return result
}

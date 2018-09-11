package leet_902

import (
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	nbytes := []byte(strconv.Itoa(N))
	setLen := len(D)
	set := make([]byte, 0, setLen)
	for _, str := range D {
		set = append(set, str[0])
	}
	ans := 0
	for i, t := 1, 1; i < len(nbytes); i++ {
		t *= setLen
		ans += t
	}
	return ans + solve(set, nbytes)
}

func solve(set []byte, nbytes []byte) int {
	if len(nbytes) == 0 {
		return 1
	}
	setLen := len(set)
	length := len(nbytes)
	idx := findIdx(set, nbytes[0])
	if idx == setLen {
		return pow(setLen, length)
	}
	if set[idx] != nbytes[0] {
		return idx * pow(setLen, length-1)
	}
	if idx == 0 {
		return solve(set, nbytes[1:])
	}
	return idx*pow(setLen, length-1) + solve(set, nbytes[1:])
}

func pow(a, b int) int {
	t := 1
	for b > 0 {
		t *= a
		b--
	}
	return t
}

func findIdx(set []byte, b byte) int {
	for idx, t := range set {
		if t >= b {
			return idx
		}
	}
	return len(set)
}

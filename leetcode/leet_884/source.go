package leet_884

import (
	"sort"
)

func decodeAtIndex(S string, K int) string {
	pos := -1
	str := []byte(S)
	var bytes [][]byte
	last := len(str) - 1
	all := 0
	var sum = []int{0}
	for pos < last {
		subStr, times := readUntil(str, &pos)
		all = (all + len(subStr)) * times
		if len(subStr) > 0 {
			sum = append(sum, all)
			bytes = append(bytes, subStr)
		} else {
			sum[len(sum)-1] = all
		}
		if all >= K {
			return string([]byte{search(sum, bytes, K)})
		}
	}
	return ""
}

// ensure that k <= sum[last]
func search(sum []int, bytes [][]byte, k int) byte {
	if k == 1 {
		return bytes[0][0]
	}
	idx := sort.SearchInts(sum, k)
	if k == sum[idx] {
		return bytes[idx-1][len(bytes[idx-1])-1]
	}
	idx--
	t := sum[idx] + len(bytes[idx])
	if t >= k {
		return bytes[idx][k-sum[idx]-1]
	}
	if k%t == 0 {
		return bytes[idx][len(bytes[idx])-1]
	}
	return search(sum, bytes, k%t)
}

func readUntil(s []byte, pos *int) ([]byte, int) {
	t := (*pos) + 1
	for i := t; i < len(s); i++ {
		if n := isNum(s[i]); n > 0 {
			*pos = i
			return s[t:i], n
		}
	}
	*pos = len(s)
	return s[t:*pos], 1
}

func isNum(b byte) int {
	if b >= '2' && b <= '9' {
		return int(b - '0')
	}
	return 0
}

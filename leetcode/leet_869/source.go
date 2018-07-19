package leet_869

func reorderedPowerOf2(N int) bool {
	mp := countMap(N)
	pre := 1
	for i := 0; i < 34; i++ {
		cur := countMap(pre)
		if mapEq(mp, cur) {
			return true
		}
		pre <<= 1
	}
	return false
}

func countMap(num int) map[int]int {
	mp := make(map[int]int)
	for num > 0 {
		t := num % 10
		if v, ok := mp[t]; ok {
			mp[t] = v + 1
		} else {
			mp[t] = 1
		}
		num /= 10
	}
	return mp
}

func mapEq(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		v2, ok := b[k]
		if !ok || v2 != v {
			return false
		}
	}
	return true
}

package leet_893

func numSpecialEquivGroups(A []string) int {
	var group []*sign
	for _, str := range A {
		s := calcSign([]byte(str))
		ok := false
		for _, g := range group {
			if cmpSign(s, g) {
				ok = true
				break
			}
		}
		if !ok {
			group = append(group, s)
		}
	}
	return len(group)
}

func cmpSign(a, b *sign) bool {
	for i := 0; i < 26; i++ {
		if a.odd[i] != b.odd[i] {
			return false
		}
		if a.even[i] != b.even[i] {
			return false
		}
	}
	return true
}

func calcSign(str []byte) *sign {
	n := len(str)
	var odd, even [26]int
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			even[str[i]-'a']++
		} else {
			odd[str[i]-'a']++
		}
	}
	return &sign{
		even: even,
		odd:  odd,
	}
}

type sign struct {
	odd  [26]int
	even [26]int
}

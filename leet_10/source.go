package leet_10

func isMatch(s string, p string) bool {
	if p == "" {
		return s == p
	}
	lenp := len(p)
	if lenp == 1 {
		return len(s) == 1 && equal(s[0], p[0])
	}
	if p[1] == '*' {
		if s == "" {
			return isMatch("", p[2:])
		}
		// match 0 p[0]
		return isMatch(s, p[2:]) ||
			// match multi p[0]
			(equal(s[0], p[0]) && isMatch(s[1:], p))
	}
	if s == "" {
		return false
	}
	return equal(s[0], p[0]) && isMatch(s[1:], p[1:])
}

func equal(a, b byte) bool {
	return a == b || b == '.'
}

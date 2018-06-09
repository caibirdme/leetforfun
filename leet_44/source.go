package leet_44

func isMatch(s string, p string) bool {
	var bytes []byte
	var pre byte
	length := len(p)
	for i := 0; i < length; i++ {
		c := p[i]
		if pre != '*' {
			bytes = append(bytes, c)
			pre = c
			continue
		}
		if c != '*' {
			bytes = append(bytes, c)
			pre = c
		}
	}
	mem := make([][]byte, length)
	lens := len(s)
	for i := 0; i < length; i++ {
		mem[i] = make([]byte, lens)
	}
	p = string(bytes)
	lenp := len(p)
	var i int
	for i = 0; i < lenp && i < lens; i++ {
		if !equal(s[i], p[i]) {
			if p[i] == '*' {
				break
			} else {
				return false
			}
		}
	}
	i = min(i, min(lenp, lens))
	return is_Match(s[i:], p[i:], mem)
}

const (
	_ byte = iota
	yes
	no
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func is_Match(s string, p string, mem [][]byte) bool {
	if p == "" && s != "" {
		return false
	}
	if p == "*" {
		return true
	}
	if s == p {
		return true
	}
	if s == "" && p != "*" {
		return false
	}
	// p != "*"
	if len(p) == 1 {
		if s == "" {
			return false
		} else if len(s) > 1 {
			return false
		}
		return equal(s[0], p[0])
	}
	lens := len(s) - 1
	lenp := len(p) - 1
	if mem[lenp][lens] == yes {
		return true
	} else if mem[lenp][lens] == no {
		return false
	}
	var res bool
	switch p[lenp] {
	case '*':
		res = is_Match(s[:lens], p, mem) || is_Match(s, p[:lenp], mem)
	case '?':
		res = is_Match(s[:lens], p[:lenp], mem)
	default:
		res = equal(s[lens], p[lenp]) && is_Match(s[:lens], p[:lenp], mem)
	}
	if res {
		mem[lenp][lens] = yes
	} else {
		mem[lenp][lens] = no
	}
	return res
}

func equal(a, b byte) bool {
	return a == b || b == '?'
}

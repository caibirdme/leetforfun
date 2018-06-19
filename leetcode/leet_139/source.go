package leet_139

func wordBreak(s string, wordDict []string) bool {
	mem := make([]byte, len(s))
	return match(s, 0, wordDict, mem)
}

const (
	_ byte = iota
	yes
	no
)

func match(s string, startPos int, wordDict []string, mem []byte) bool {
	lens := len(s)
	if startPos >= lens {
		return true
	}
	if mem[startPos] == yes {
		return true
	} else if mem[startPos] == no {
		return false
	}
	p := s[startPos]
	for i := 0; i < len(wordDict); i++ {
		lenwd := len(wordDict[i])
		endPos := startPos + lenwd
		if wordDict[i][0] != p || endPos > lens {
			continue
		}
		t := s[startPos:endPos]
		if t == wordDict[i] && match(s, endPos, wordDict, mem) {
			mem[startPos] = yes
			return true
		}
	}
	mem[startPos] = no
	return false
}

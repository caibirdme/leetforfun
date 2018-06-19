package leet_132

func minCut(s string) int {
	lens := len(s)
	if lens <= 1 {
		return 0
	}
	palindome := make([][]byte, lens)
	mem := make([][]int, lens)
	for i := 0; i < lens; i++ {
		palindome[i] = make([]byte, lens)
		mem[i] = make([]int, lens)
	}
	return search([]byte(s), 0, lens-1, mem, palindome)
}

const (
	undefine byte = iota
	yes
	no
)

func search(s []byte, left, right int, mem [][]int, palindome [][]byte) int {
	if isPalindome(s, left, right, palindome) == yes {
		return 0
	}
	if mem[left][right] > 0 {
		return mem[left][right]
	}
	mem[left][right] = right - left
	for i := right - 1; i >= left; i-- {
		if isPalindome(s, left, i, palindome) != yes {
			continue
		}
		t := search(s, i+1, right, mem, palindome)
		mem[left][right] = min(mem[left][right], t+1)
	}
	return mem[left][right]
}

func isPalindome(s []byte, left, right int, palindome [][]byte) byte {
	if left == right {
		return yes
	}
	if palindome[left][right] != undefine {
		return palindome[left][right]
	}
	if s[left] != s[right] {
		palindome[left][right] = no
		return no
	}
	if left+1 == right {
		palindome[left][right] = yes
		return yes
	}
	palindome[left][right] = isPalindome(s, left+1, right-1, palindome)
	return palindome[left][right]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

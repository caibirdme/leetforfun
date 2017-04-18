package leet_22

var res []string

func generateParenthesis(n int) []string {
	res = nil
	gen(n, n, 0, make([]byte, n<<1))
	return res
}

func gen(left, right, idx int, arr []byte) {
	if idx == cap(arr) {
		res = append(res, string(arr))
		return
	}
	if left > 0 {
		arr[idx] = '('
		gen(left-1, right, idx+1, arr)
	}
	if right > left {
		arr[idx] = ')'
		gen(left, right-1, idx+1, arr)
	}
}

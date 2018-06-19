package leet_17

var (
	answer [][]byte

	mapping = map[byte][]byte{
		'0': []byte{' '},
		'1': []byte{0},
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
)

func letterCombinations(digits string) []string {
	if "" == digits {
		return nil
	}
	answer = nil
	combine([]byte(digits), make([]byte, len(digits)), 0)
	var res []string
	for _, item := range answer {
		res = append(res, string(item))
	}
	return res
}

func combine(str []byte, cur []byte, idx int) {
	if nil == str {
		arr := make([]byte, idx)
		copy(arr, cur)
		answer = append(answer, arr)
		return
	}
	if str[0] == '1' {
		if len(str) == 1 {
			combine(nil, cur, idx)
		} else {
			combine(str[1:], cur, idx)
		}
		return
	}
	bytes := mapping[str[0]]
	for _, c := range bytes {
		cur[idx] = c
		if len(str) == 1 {
			combine(nil, cur, idx+1)
		} else {
			combine(str[1:], cur, idx+1)
		}
	}
}

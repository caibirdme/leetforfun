package leet_93

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	var ans []string
	data := make([]string, 4)
	dfs(0, data, []byte(s), &ans)
	return ans
}

func dfs(idx int, data []string, str []byte, ans *[]string) {
	if idx == 3 {
		if isValid(str) {
			data[idx] = string(str)
			*ans = append(*ans, strings.Join(data, "."))
		}
		return
	}
	length := len(str)
	if length == 0 {
		return
	}
	data[idx] = string(str[0])
	dfs(idx+1, data, str[1:], ans)
	if length >= 2 && isValid(str[:2]) {
		data[idx] = string(str[:2])
		dfs(idx+1, data, str[2:], ans)
	}
	if length >= 3 && isValid(str[:3]) {
		data[idx] = string(str[:3])
		dfs(idx+1, data, str[3:], ans)
	}
}

func isValid(str []byte) bool {
	length := len(str)
	if length > 3 || length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	if length == 2 {
		return str[0] != '0'
	}
	if str[0] > '2' || str[0] == '0' {
		return false
	}
	if str[0] == '1' {
		return true
	}
	if str[1] > '5' || (str[1] == '5' && str[2] > '5') {
		return false
	}
	return true
}

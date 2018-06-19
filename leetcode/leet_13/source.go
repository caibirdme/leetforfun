package leet_13

import (
	"bytes"
)

var order = [7]byte{'M','D','C','L','X','V','I'}
var value = [7]int{1000,500,100,50,10,5,1}

func romanToInt(s string) int {
	if "" == s {return 0}
	return convert([]byte(s))
}

func convert(num []byte) int {
	if nil == num || 0 == len(num) {return 0}
	var idx int
	for i,c := range order {
		idx = bytes.IndexByte(num, c)
		if -1 == idx {continue}
		if idx != 0 {
			return value[i]-convert(num[:idx])+convert(num[idx+1:])
		}
		return value[i]+convert(num[1:])
	}
	return 0
}
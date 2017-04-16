package leet_6

import (
	"bytes"
)

type byteSlice []byte

func convert(s string, numRows int) string {
	if numRows == 1 || s == "" {
		return s
	}
	arr := make([]byteSlice, numRows)
	length := len(s)
	period := (numRows - 1) << 1
	for i := 0; i < length; i += period {
		var j, k int
		for j = 0; j < numRows && i+j < length; j++ {
			arr[j] = append(arr[j], s[i+j])
		}
		for j, k = numRows-2, numRows; j > 0 && i+k < length; j, k = j-1, k+1 {
			arr[j] = append(arr[j], s[i+k])
		}
	}
	var b bytes.Buffer
	for _, item := range arr {
		b.Write(item)
	}
	return b.String()
}

package leet_68

import (
	"bytes"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	for len(words) > 0 {
		pos, str := peekWords(words, maxWidth)
		ans = append(ans, str)
		words = words[pos:]
	}
	return ans
}

func peekWords(words []string, maxWidth int) (int, string) {
	count := len(words[0])
	if count == maxWidth {
		return 1, words[0]
	}
	num := 1
	for i := 1; i < len(words); i++ {
		t := count + num + len(words[i])
		if t > maxWidth {
			break
		}
		count += len(words[i])
		num++
	}
	var buf = bytes.NewBuffer(nil)
	if num == len(words) {
		buf.WriteString(words[0])
		for i := 1; i < num; i++ {
			buf.WriteString(" ")
			buf.WriteString(words[i])
		}
		extraSpace := maxWidth - count - (num - 1)
		for i := 0; i < extraSpace; i++ {
			buf.WriteString(" ")
		}
		return num, buf.String()
	}
	buf.WriteString(words[0])
	space := maxWidth - count
	if num == 1 {
		buf.WriteString(strings.Repeat(" ", space))
		return 1, buf.String()
	}
	perSpace := space / (num - 1)
	extraSpace := space % (num - 1)
	for i := 1; i < num; i++ {
		t := perSpace
		if extraSpace > 0 {
			t++
			extraSpace--
		}
		buf.WriteString(strings.Repeat(" ", t))
		buf.WriteString(words[i])
	}
	return num, buf.String()
}

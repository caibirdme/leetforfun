package leet_72

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	var testData = []struct {
		word1 string
		word2 string
		out   int
	}{
		{
			word1: "inten",
			word2: "execu",
			out:   5,
		},
		{
			word1: "intention",
			word2: "execution",
			out:   5,
		},
		{
			word1: "abcef",
			word2: "defgef",
			out:   4,
		},
		{
			word1: "cbaaf",
			word2: "cbaadf",
			out:   1,
		},
		{
			word1: "cbaaf",
			word2: "cdaadf",
			out:   2,
		},
		{
			word1: "aaeacbaaf",
			word2: "aacacdaadf",
			out:   3,
		},
		{
			word1: "",
			word2: "bbabba",
			out:   6,
		},
		{
			word1: "",
			word2: "",
			out:   0,
		},
		{
			word1: "a",
			word2: "bbabba",
			out:   5,
		},
		{
			word1: "bbabba",
			word2: "a",
			out:   5,
		},
		{
			word1: "abcacbaa",
			word2: "abcacaa",
			out:   1,
		},
		{
			word1: "abcacbaa",
			word2: "abcacaad",
			out:   2,
		},
		{
			word1: "abcacbaaf",
			word2: "abcacaadf",
			out:   2,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, minDistance(tc.word1, tc.word2), "word1=%s\r\nword2=%s\r\n", tc.word1, tc.word2)
	}
}

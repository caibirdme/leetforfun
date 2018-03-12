package leet_516

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	var testData = []struct {
		in  string
		out int
	}{
		{
			in:  "",
			out: 0,
		},
		{
			in:  "a",
			out: 1,
		},
		{
			in:  "ab",
			out: 1,
		},
		{
			in:  "aba",
			out: 3,
		},
		{
			in:  "cbbd",
			out: 2,
		},
		{
			in:  "bbbab",
			out: 4,
		},
		{
			in:  "aaqbbchcdmdeejddcwctbsbaar",
			out: 18,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, longestPalindromeSubseq(tc.in), "%s", tc.in)
	}
}

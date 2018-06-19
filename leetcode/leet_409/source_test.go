package leet_409

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
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
			in:  "abccccdd",
			out: 7,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, longestPalindrome(tc.in), "%s", tc.in)
	}
}

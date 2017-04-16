package leet_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type tCase struct {
	str, ans string
}

func TestLongestPalindrome(t *testing.T) {
	var data = []tCase{
		tCase{
			"babad",
			"bab",
		},
		tCase{
			"cbbd",
			"bb",
		},
		tCase{
			"acbaabcbaa",
			"aabcbaa",
		},
		tCase{
			"abcdefg",
			"a",
		},
		tCase{
			"abcdefghijklkl",
			"klk",
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.ans, longestPalindrome(tc.str))
	}
}

package leet_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	s   string
	ans int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var data = []testCase{
		testCase{
			"abcabcbb",
			3,
		},
		testCase{
			"bbbbbb",
			1,
		},
		testCase{
			"pwwkew",
			3,
		},
		testCase{
			"tmmzuxt",
			5,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.ans, lengthOfLongestSubstring(tc.s))
	}
}

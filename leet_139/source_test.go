package leet_139

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordBreak(t *testing.T) {
	var testCase = []struct {
		s        string
		wordDict []string
		expect   bool
	}{
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expect:   true,
		},
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expect:   true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expect:   false,
		},
		{
			s:        "catsandsandcatsandcatdogdogcatcatsandcatscatdogsand",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expect:   true,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := wordBreak(tc.s, tc.wordDict)
		ass.Equal(tc.expect, actual, "s: %s dict: %v", tc.s, tc.wordDict)
	}
}

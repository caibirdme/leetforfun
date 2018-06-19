package leet_140

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordBreak(t *testing.T) {
	var testCase = []struct {
		s        string
		wordDict []string
		expect   []string
	}{
		{
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			expect: []string{
				"cats and dog",
				"cat sand dog",
			},
		},
		{
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			expect: []string{
				"pine apple pen apple",
				"pineapple pen apple",
				"pine applepen apple",
			},
		},
		{
			s:        "foobar",
			wordDict: []string{"foo", "bar", "fo", "o", "f"},
			expect: []string{
				"foo bar",
				"fo o bar",
				"f o o bar",
			},
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expect:   []string{},
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := wordBreak(tc.s, tc.wordDict)
		sort.Strings(actual)
		sort.Strings(tc.expect)
		ass.Equal(tc.expect, actual, "s: %s dict: %v", tc.s, tc.wordDict)
	}
}

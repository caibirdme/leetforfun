package leet_10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsMatch(t *testing.T) {
	var testCase = []struct {
		s       string
		pattern string
		expect  bool
	}{
		{
			s:       "aa",
			pattern: "a",
			expect:  false,
		},
		{
			s:       "aa",
			pattern: "aa",
			expect:  true,
		},
		{
			s:       "aa",
			pattern: "a*",
			expect:  true,
		},
		{
			s:       "aa",
			pattern: ".*",
			expect:  true,
		},
		{
			s:       "aab",
			pattern: "c*a*b",
			expect:  true,
		},
		{
			s:       "my name is deen",
			pattern: "my name is .*",
			expect:  true,
		},
		{
			s:       "my name is ",
			pattern: "my name is .*",
			expect:  true,
		},
		{
			s:       "my name is            ajsdhakshdjkahsd",
			pattern: "my name is .*",
			expect:  true,
		},
		{
			s:       "my name is",
			pattern: "my name is .*",
			expect:  false,
		},
		{
			s:       "I love playing soccer",
			pattern: "I .*soccer",
			expect:  true,
		},
		{
			s:       "I hate playing soccer",
			pattern: "I .*soccer",
			expect:  true,
		},
		{
			s:       "I soccer",
			pattern: "I .*soccer",
			expect:  true,
		},
		{
			s:       "I   soccer",
			pattern: "I .*soccer",
			expect:  true,
		},
		{
			s:       "I have a lot of soccers",
			pattern: "I .*soccer",
			expect:  false,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		ass.Equal(tc.expect, isMatch(tc.s, tc.pattern), "s: %s p: %s", tc.s, tc.pattern)
	}
}

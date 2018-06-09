package leet_44

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsMatch(t *testing.T) {
	var testCase = []struct {
		s      string
		p      string
		expect bool
	}{
		{
			s:      "adcebadcebadcebadcebadcebadcebadcebadcebadcebadcebadcebghtasdsadkkasdas",
			p:      "adcebadcebadcebadcebadcebadcebadcebadcebadcebadcebadreb*?t*sad?*asda?",
			expect: false,
		},
		{
			s:      "adcebadcebadcebadcebadcebadcebadcebadcebadcebadcebadcebghtasdsadkkasdas",
			p:      "adcebadcebadcebadcebadcebadcebadcebadcebadcebadcebadceb*?t*sad?*asda?",
			expect: true,
		},
		{
			s:      "adceb",
			p:      "*a*b",
			expect: true,
		},
		{
			s:      "ab",
			p:      "ab**",
			expect: true,
		},
		{
			s:      "acdcb",
			p:      "a*c?b",
			expect: false,
		},
		{
			s:      "aa",
			p:      "a?",
			expect: true,
		},
		{
			s:      "aa",
			p:      "a",
			expect: false,
		},
		{
			s:      "aasdasdaa",
			p:      "",
			expect: false,
		},
		{
			s:      "aa",
			p:      "*",
			expect: true,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := isMatch(tc.s, tc.p)
		ass.Equal(tc.expect, actual, "s: %s p: %s", tc.s, tc.p)
	}
}

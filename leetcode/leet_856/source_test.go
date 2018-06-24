package leet_856

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScoreOfParentheses(t *testing.T) {
	var testCase = []struct {
		s   string
		exp int
	}{
		{
			s:   "()()",
			exp: 2,
		},
		{
			s:   "()",
			exp: 1,
		},
		{
			s:   "(())",
			exp: 2,
		},
		{
			s:   "(()())",
			exp: 4,
		},
		{
			s:   "(()(()))",
			exp: 6,
		},
		{
			s:   "()((()(())()))(())",
			exp: 19,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := scoreOfParentheses(tc.s)
		should.Equal(tc.exp, actual, "S: %s", tc.s)
	}
}

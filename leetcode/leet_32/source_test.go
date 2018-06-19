package leet_32

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  string
	out int
}

func TestLongestValidParentheses(t *testing.T) {
	var data = []Case{
		Case{
			")()())",
			4,
		},
		Case{
			"()())(()())",
			6,
		},
		Case{
			"",
			0,
		},
		Case{
			"(",
			0,
		},
		Case{
			"()",
			2,
		},
		Case{
			"()(()())",
			8,
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, longestValidParentheses(tc.in), "in: %s", tc.in)
	}
}

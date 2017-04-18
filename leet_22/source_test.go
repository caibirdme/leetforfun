package leet_22

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Case struct {
	in  int
	out []string
}

func TestGenerateParenthesis(t *testing.T) {
	var data = []Case{
		Case{
			3,
			[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		Case{
			2,
			[]string{
				"()()",
				"(())",
			},
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		sort.Strings(tc.out)
		ans := generateParenthesis(tc.in)
		sort.Strings(ans)
		ass.Equal(tc.out, ans, "in: %d", tc.in)
	}
}

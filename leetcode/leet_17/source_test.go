package leet_17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  string
	out []string
}

func TestLetterCombinations(t *testing.T) {
	var data = []Case{
		Case{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		Case{
			"10",
			[]string{" "},
		},
		Case{
			"1",
			[]string{""},
		},
		Case{
			"",
			nil,
		},
		Case{
			"9",
			[]string{"w", "x", "y", "z"},
		},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, letterCombinations(tc.in), "in: %s", tc.in)
	}
}

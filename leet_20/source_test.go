package leet_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  string
	out bool
}

func TestIsValid(t *testing.T) {
	var data = []Case{
		Case{"]", false},
		Case{"()", true},
		Case{"()[]{}", true},
		Case{"([)]", false},
		Case{"((({{}}[[]])))", true},
		Case{"((({{}}[[]]))", false},
	}
	ass := assert.New(t)
	for _, tc := range data {
		ass.Equal(tc.out, isValid(tc.in), "in: %s", tc.in)
	}
}

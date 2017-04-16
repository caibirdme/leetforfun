package leet_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Case struct {
	in  string
	out int
}

func TestMyAtoi(t *testing.T) {
	var data = []Case{
		Case{"", 0},
		Case{"     ", 0},
		Case{"0     ", 0},
		Case{"0-0123", -123},
		Case{"+123", 123},
		Case{"-0012", -12},
		Case{"    -0012a42", -12},
		Case{"   +12 3 4 5      ", 12},
		Case{"   +0 123", 0},
		Case{"2147483648", 2147483647},
		Case{"-2147483649", -2147483648},
	}
	ass := assert.New(t)
	for _, tc := range data {
		out, err := convert(tc.in)
		ass.NoError(err, "in: %s, out:%d", tc.in, tc.out)
		ass.Equal(tc.out, out)
	}
}
